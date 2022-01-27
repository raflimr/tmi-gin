package api

import (
	"log"
	"net/http"
	"strconv"

	db "tmi-gin/db/sqlc"

	"github.com/gin-gonic/gin"
)

type listPracticeByCategory struct {
	IDCategory int32 `form:"id_category"`
	PageID     int32 `form:"page_id" binding:"required,min=1"`
	PageSize   int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPracticesByCategory(ctx *gin.Context) {
	idCategory := ctx.Param("id_category")
	idCategoryInt, _ := strconv.Atoi(idCategory)
	var req listPracticeByCategory
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPracticeByCategoryParams{
		IDCategory: int32(idCategoryInt),
		Limit:      req.PageSize,
		Offset:     (req.PageID - 1) * req.PageSize,
	}
	log.Println(arg)

	practice, err := server.store.ListPracticeByCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, practice)
}

type listPracticeRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPractices(ctx *gin.Context) {
	var req listPracticeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPracticeParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	log.Println(arg)

	practice, err := server.store.ListPractice(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, practice)
}

type practiceDetailRequest struct {
	ID int32 `form:"id_practice"`
}

func (server *Server) practiceDetailStatistik(ctx *gin.Context) {
	var req practiceDetailRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	practice, err := server.store.InfoPracticeStatistik(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, practice)
}

func (server *Server) practiceDetail(ctx *gin.Context) {
	var req practiceDetailRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	practice, err := server.store.InfoPractice(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, practice)
}

func (server *Server) practiceVideoById(ctx *gin.Context) {
	var req practiceDetailRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	practice, err := server.store.VideoInPractice(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, practice)
}
