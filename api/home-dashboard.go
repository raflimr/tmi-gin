package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	db "tmi-gin/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (server *Server) homeDashboard(ctx *gin.Context) {
	id := ctx.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	mahasiswa, err := server.store.HomeDashboard(ctx, int32(idInt))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var validation db.HomeDashboardRow

	if mahasiswa.NamaLengkap.Valid {
		validation.NamaLengkap = mahasiswa.NamaLengkap
	} else {
		validation.NamaLengkap.Valid = false
	}

	if mahasiswa.UrlFoto.Valid {
		validation.UrlFoto = mahasiswa.UrlFoto
	} else {
		validation.UrlFoto.String = "Null"
		fmt.Println(validation.UrlFoto.String)
	}

	if mahasiswa.University.Valid {
		validation.University = mahasiswa.University
	} else {
		validation.University.Valid = false
	}

	if mahasiswa.FinishedPractice.Valid {
		validation.FinishedPractice = mahasiswa.FinishedPractice
	} else {
		validation.FinishedPractice.Valid = false
	}

	if mahasiswa.FinishedChallange.Valid {
		validation.FinishedPractice = mahasiswa.FinishedPractice
	} else {
		validation.FinishedPractice.Valid = false
	}

	if mahasiswa.TotalSkor.Valid {
		validation.TotalSkor = mahasiswa.TotalSkor
	} else {
		validation.TotalSkor.Valid = false
	}

	rsp := db.HomeDashboardRow{
		Username:          mahasiswa.Username,
		NamaLengkap:       validation.NamaLengkap,
		UrlFoto:           validation.UrlFoto,
		University:        validation.University,
		FinishedPractice:  validation.FinishedPractice,
		JumlahPractice:    mahasiswa.JumlahPractice,
		FinishedChallange: validation.FinishedChallange,
		JumlahChallange:   mahasiswa.JumlahChallange,
		TotalSkor:         validation.TotalSkor,
	}

	ctx.JSON(http.StatusOK, rsp)
}
