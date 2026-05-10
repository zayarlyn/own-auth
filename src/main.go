package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zayarlyn/own-auth/src/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializeDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	db := initializeDB()

	router.Register(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
