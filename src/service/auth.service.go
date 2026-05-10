// Package service
package service

import (
	"context"
	"log"

	"github.com/zayarlyn/own-auth/src/db/model"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) string {
	user, err := gorm.G[model.User](db).First(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return user.ID
}
