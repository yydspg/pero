package core

import (
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PeroDB struct {
	db  *gorm.DB
	log *zerolog.Logger
}

var peroDB *PeroDB

func init() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	peroDB.db = db
	if err != nil {
		peroDB.log.Panic().Err(err).Msg("failed to connect to database")
		panic(err)
	}
}
func getPeroDB() *PeroDB {
	return peroDB
}
