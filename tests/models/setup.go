package models_test

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/config"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/db"
)

var TRdb *db.Rdb

func SetupTestDB(appConfig *config.Config) error {
	rdb, err := db.New(mysql.Config{
		User:                 appConfig.MysqlUser,
		Passwd:               appConfig.MysqlPassword,
		Addr:                 appConfig.MysqlAddr,
		DBName:               appConfig.MysqlDatabase,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	TRdb = rdb
	if err != nil {
		log.Fatal(err)
	}
	TRdb.Migration()
	return nil
}

func TeardownTestDB() {
	// cleanup
	var tableNames []string
	TRdb.Db.Raw("SHOW TABLES").Scan(&tableNames)
	// Iterate over each table name and drop it
	for _, tableName := range tableNames {
		TRdb.Db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName))
	}
}
