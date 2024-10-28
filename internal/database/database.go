package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Conn struct {
	db *gorm.DB
}

func ConnectDatabase() Conn {
	var conn Conn
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		panic(err)
	}
	conn.db = db
	return conn
}
