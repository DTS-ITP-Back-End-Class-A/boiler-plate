package configs

import (
	"fmt"
	"log"
	"os"
	"testing/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func InitDB(cfg models.Config) {
	var err error
	// mysql
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	dbConn, err = gorm.Open(mysql.Open(connStr), &gorm.Config{
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Printf("Error in creating the new connection: %v.", err)
		os.Exit(1)
		return
	}
}
