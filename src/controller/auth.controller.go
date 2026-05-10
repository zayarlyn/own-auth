// Package controller
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zayarlyn/own-auth/src/service"
	"gorm.io/gorm"
)

func Login(c *gin.Context, db *gorm.DB) {
	jwt := service.Login(db)

	c.JSON(200, gin.H{
		"accessToken": jwt,
	})
}
