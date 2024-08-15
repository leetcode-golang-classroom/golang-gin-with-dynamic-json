package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/config"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/db"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/util"
)

type App struct {
	router *gin.Engine
	config *config.Config
	rdb    *db.Rdb
}

func New(config *config.Config) *App {
	rdb, err := db.New(mysql.Config{
		User:                 config.MysqlUser,
		Passwd:               config.MysqlPassword,
		Addr:                 config.MysqlAddr,
		DBName:               config.MysqlDatabase,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	rdb.Migration()
	app := &App{
		config: config,
		rdb:    rdb,
	}
	app.loadRoutes()
	app.SetUpBlogsRoutes()
	return app
}

func (app *App) Start(ctx context.Context) error {
	// setup listener with port
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.Port),
		Handler: app.router,
	}
	log.Printf("Starting server on %s", app.config.Port)
	errCh := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			errCh <- fmt.Errorf("failed to start server: %w", err)
		}
		util.CloseChannel(errCh)
	}()
	defer func() {
		log.Println("before tear down app")
	}()
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
