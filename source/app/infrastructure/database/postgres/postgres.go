package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PW"), os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_DB"))

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
