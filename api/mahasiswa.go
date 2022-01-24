package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	db "tmi-gin/db/sqlc"
	"tmi-gin/util"

	"github.com/gin-gonic/gin"
)

type createMahasiswaRequest struct {
	Username string `json:"username" binding:"alphanum"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"email"`
	Error    error  `json:"error"`
}

type mahasiswaResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(m db.Mahasiswa) mahasiswaResponse {
	return mahasiswaResponse{
		Username:  m.Username,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
	}
}

func (server *Server) createMahasiswa(ctx *gin.Context) {
	var req createMahasiswaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, req.Error)
		return
	}

	username, err := server.store.CheckUsername(context.TODO(), req.Username)
	if err == nil && username.Username == req.Username {
		ctx.JSON(http.StatusNotAcceptable, "error : Username sudah di pake")
		return
	}

	email, err := server.store.CheckEmail(context.TODO(), req.Email)
	if err == nil && email.Email == req.Email {
		ctx.JSON(http.StatusNotAcceptable, "error : Email sudah di pake")
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateMahasiswaParams{
		Username:      req.Username,
		Password:      hashedPassword,
		Email:         req.Email,
		NomorHp:       "",
		UrlFoto:       "",
		NamaLengkap:   "",
		TanggalLahir:  "",
		JenisKelamin:  "",
		University:    "",
		Nim:           "",
		TahunMasuk:    "",
		KotaKabupaten: "",
		Token:         "",
	}

	mahasiswa, err := server.store.CreateMahasiswa(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(mahasiswa)

	ctx.JSON(http.StatusOK, rsp)
}

type loginMahasiswaRequest struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email"`
}

type loginMahasiswaResponse struct {
	AccessToken string            `json:"access_token"`
	Mahasiswa   mahasiswaResponse `json:"mahasiswa"`
}

func (server *Server) loginMahasiswa(ctx *gin.Context) {
	var req loginMahasiswaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetMahasiswaParams{
		Username: req.Username,
		Email:    req.Email,
	}

	mahasiswa, err := server.store.GetMahasiswa(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	log.Println(mahasiswa.Password)

	verify, err := util.CheckPassword(req.Password, mahasiswa.Password)
	log.Println(verify)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		mahasiswa.Username,
		5000000000,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginMahasiswaResponse{
		AccessToken: accessToken,
		Mahasiswa:   newUserResponse(mahasiswa),
	}

	ctx.JSON(http.StatusAccepted, rsp)
}

// q := url.Values{}
// id := strconv.Itoa(int(mahasiswa.ID))
// q.Set("id", id)
// location := url.URL{Path: "/users/home-dashboard", RawQuery: q.Encode()}
// ctx.Redirect(http.StatusNotFound, location.RequestURI())
//PUT
type updateMahasiswaResponse struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	NomorHp       string `json:"nomor_hp"`
	Email         string `json:"email"`
	UrlFoto       string `json:"url_foto"`
	NamaLengkap   string `json:"nama_lengkap"`
	TanggalLahir  string `json:"tanggal_lahir"`
	JenisKelamin  string `json:"jenis_kelamin"`
	University    string `json:"university"`
	Nim           string `json:"nim"`
	Jurusan       string `json:"jurusan"`
	TahunMasuk    string `json:"tahun_masuk"`
	KotaKabupaten string `json:"kota_kabupaten"`
}

func (server *Server) updateMahasiswa(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	var req db.UpdateMahasiswaParams
	req.ID = int32(idInt)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	rsp := updateMahasiswaResponse{
		Username:      req.Username,
		Password:      req.Password,
		NomorHp:       req.NomorHp,
		Email:         req.Email,
		UrlFoto:       req.UrlFoto,
		NamaLengkap:   req.NamaLengkap,
		TanggalLahir:  req.TanggalLahir,
		JenisKelamin:  req.JenisKelamin,
		University:    req.University,
		Nim:           req.Nim,
		Jurusan:       req.Jurusan,
		TahunMasuk:    req.TahunMasuk,
		KotaKabupaten: req.KotaKabupaten,
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	req.Password = string(hashedPassword)

	if err := server.store.UpdateMahasiswa(ctx, req); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

var otp *string

type postEmailOTP struct {
	Email string `json:"email"`
}

func (server *Server) postEmailOtp(ctx *gin.Context) {
	var req postEmailOTP
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	email, err := server.store.CheckEmail(context.TODO(), req.Email)
	if err == nil && email.Email == req.Email {
		otp, err = util.OTPEmail(req.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		ctx.JSON(http.StatusAccepted, "Email terkirim")
		return
	}
}

func (server *Server) updateOTP(ctx *gin.Context) {
	var req postEmailOTP
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateOTPInDBParams{
		Email: req.Email,
		Token: *otp,
	}

	err := server.store.UpdateOTPInDB(context.TODO(), arg)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, arg)
}

type tokenVerifikasi struct {
	Token string `json:"token" binding:"numeric"`
}

func (server *Server) TokenVerifikasi(ctx *gin.Context) {
	var req tokenVerifikasi
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	token, err := server.store.CheckToken(ctx, req.Token)
	switch {
	case token.Token == req.Token && err == nil:
		ctx.JSON(http.StatusAccepted, "Token berhasil diinput")
		return
	default:
		ctx.JSON(http.StatusNotAcceptable, "Token Expired")
		return
	}
}

type changePassword struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (server *Server) changePassword(ctx *gin.Context) {
	var req changePassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ChangePasswordParams{
		Email:    req.Email,
		Password: req.Password,
	}

	err := server.store.ChangePassword(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Input yang benar")
	}

	ctx.JSON(http.StatusAccepted, arg)

}
