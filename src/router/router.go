// Package router
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zayarlyn/own-auth/src/controller"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB) {
	authRouter := r.Group("/auth")

	authRouter.GET("/login", func(ctx *gin.Context) {
		controller.Login(ctx, db)
	})
}
