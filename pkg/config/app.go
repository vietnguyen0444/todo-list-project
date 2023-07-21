package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user:sa password:1234567 dbname:GoGormORM port:7071 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDatabase() *gorm.DB {
	return db
}
