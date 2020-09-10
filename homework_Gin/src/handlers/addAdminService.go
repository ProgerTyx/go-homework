package handler

import (
	"github.com/gin-gonic/gin"
	"go-homework/homework_Gin/config"
	models "go-homework/homework_Gin/src/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(ctx *gin.Context) {
	admin := &models.Admin{Login: "admin", Password: "admin"}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	admin.Password = string(hashedPassword)
	db, _ := config.GetMongoDB()
	_ = db.C("admin").Insert(admin)

	ctx.JSON(200, gin.H{
		"message": "Admin Created",
	})
}
