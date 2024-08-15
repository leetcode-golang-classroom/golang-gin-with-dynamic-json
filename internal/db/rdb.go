package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Rdb struct {
	Db *gorm.DB
}

func New(cfg mysql.Config) (*Rdb, error) {
	log.Printf("%v\n", cfg)
	db, err := gorm.Open(gorm_mysql.Open(cfg.FormatDSN()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect db: %s", err)
		return nil, err
	}
	return &Rdb{Db: db}, nil
}
