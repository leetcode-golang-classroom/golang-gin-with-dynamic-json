package blog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/types"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/util"
)

type Handler struct {
	store types.BlogRepo
}

func NewHandler(store types.BlogRepo) *Handler {
	return &Handler{
		store: store,
	}
}

func (handler *Handler) RegisterRoute(router *gin.RouterGroup) {
	router.GET("/", handler.BlogIndex)
	router.GET("/:id", handler.BlogShow)
}

func (handler *Handler) BlogIndex(ctx *gin.Context) {
	blogs, err := handler.store.FindAll(ctx)
	if err != nil {
		util.WriteError(ctx.Writer, http.StatusInternalServerError, err)
		return
	}
	switch getFormat(ctx) {
	case formatHTML:
		ctx.HTML(http.StatusOK,
			"blog/index.tpl",
			gin.H{
				"blogs":      blogs,
				"page":       ctx.GetInt("page"),
				"pageSize":   ctx.GetInt("pageSize"),
				"totalPages": ctx.GetInt("totalPages"),
			},
		)
	case formatJSON:
		ctx.JSON(http.StatusOK, gin.H{
			"blogs":      blogs,
			"page":       ctx.GetInt("page"),
			"pageSize":   ctx.GetInt("pageSize"),
			"totalPages": ctx.GetInt("totalPages"),
		})
	}
}

func (handler *Handler) BlogShow(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		util.WriteError(ctx.Writer, http.StatusBadRequest, err)
		return
	}
	blog, err := handler.store.Find(ctx, id)
	if err != nil {
		util.WriteError(ctx.Writer, http.StatusInternalServerError, err)
		return
	}
	switch getFormat(ctx) {
	case formatHTML:
		ctx.HTML(http.StatusOK,
			"blog/show.tpl",
			gin.H{
				"blog": blog,
			},
		)
	case formatJSON:
		ctx.JSON(http.StatusOK,
			gin.H{
				"blog": blog,
			},
		)
	}
}
