package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/service/blog"
)

// define route
func (app *App) loadRoutes() {
	gin.SetMode(app.config.GinMode)
	router := gin.New()
	// recovery middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	// setup template
	router.LoadHTMLGlob("template/**/*")
	// default health
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "status ok"})
	})
	app.router = router
}

func (app *App) SetUpBlogsRoutes() {
	blogsGroup := app.router.Group("/blogs")
	blogsStore := blog.NewBlogStore(app.rdb)
	blogsHdr := blog.NewHandler(blogsStore)
	blogsHdr.RegisterRoute(blogsGroup)
}
