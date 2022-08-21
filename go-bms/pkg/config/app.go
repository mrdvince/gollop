package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connect() *gorm.DB {
	//single variable for everything because am lazy
	dsn := os.ExpandEnv("host=localhost user=$GORM password=$GORM dbname=$GORM port=5432 sslmode=disable TimeZone=Europe/Amsterdam")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func GetDB() *gorm.DB {
	return Connect()
}
