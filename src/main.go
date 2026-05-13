package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zayarlyn/own-auth/src/db/model"
	"github.com/zayarlyn/own-auth/src/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializeAndMigrateDB() *gorm.DB {
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s`, os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"), os.Getenv("DB_PUB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if migrationErr := db.AutoMigrate(&model.User{}, &model.UserPool{}); migrationErr != nil {
		log.Fatal(migrationErr)
	}

	return db
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	db := initializeAndMigrateDB()

	router.Register(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
