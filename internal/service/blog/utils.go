package blog

import "github.com/gin-gonic/gin"

type respFormat int

const (
	formatHTML respFormat = iota
	formatJSON
)

var formatMap = map[string]respFormat{
	"json": formatJSON,
	"html": formatHTML,
}

func getFormat(ctx *gin.Context) respFormat {
	format := ctx.DefaultQuery("format", "html")
	return formatMap[format]
}
