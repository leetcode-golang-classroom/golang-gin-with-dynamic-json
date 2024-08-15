package types

import (
	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/model"
)

type BlogRepo interface {
	FindAll(ctx *gin.Context) (*[]model.Blog, error)
	Find(ctx *gin.Context, id uint64) (*model.Blog, error)
}
