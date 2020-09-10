package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-homework/homework_Gin/src/handlers"
	models "go-homework/homework_Gin/src/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

const adminCollection = "admin"

func Login(ctx *gin.Context) string {
	db := *handler.MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	var loginData models.Admin
	err := ctx.BindJSON(&loginData)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return ""
	}

	var admin models.Admin
	err = db.C(adminCollection).Find(bson.M{"login": loginData.Login}).One(&admin)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error Find Login",
		})
		return ""
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginData.Password))
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error Decoded Password",
		})
		return ""
	}

	return JWTAuthService().GenerateToken()
}
