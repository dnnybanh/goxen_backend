package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	userRoutes := r.Group("/user")
	userRoutes.POST("/register", h.CreateUser)
}