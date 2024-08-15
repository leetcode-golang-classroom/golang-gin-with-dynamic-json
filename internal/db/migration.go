package db

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/model"
)

func (rdb *Rdb) Migration() {
	log.Printf("start migration")
	rdb.Db.AutoMigrate(&model.Blog{})
}
