package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-homework/homework_Gin/config"
	models "go-homework/homework_Gin/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const userCollection = "users"

func MongoConfig() *mgo.Database {
	db, err := config.GetMongoDB()
	if err != nil {
		panic(err)
	}
	return db
}

func GetAllUser(ctx *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	var users models.Users
	err := db.C(userCollection).Find(bson.M{}).All(&users)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error Get All User",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"user": &users,
	})
}

func GetUser(ctx *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	id := ctx.Param("id")

	var user models.User
	err := db.C(userCollection).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "User not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"user": &user,
	})
}

func CreateUser(ctx *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	var user models.User
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return
	}

	err = db.C(userCollection).Insert(user)

	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error Insert User",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success Insert User",
		"user":    &user,
	})
}

func UpdateUser(ctx *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	id := ctx.Param("id")

	var user models.User
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return
	}
	change := bson.M{"$set": bson.M{"firstname": user.FirstName, "lastname": user.LastName}}
	err = db.C(userCollection).Update(bson.M{"_id": bson.ObjectIdHex(id)}, change)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error Update User",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success Update User",
		"user":    &user,
	})
}

func DeleteUser(ctx *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	id := ctx.Param("id")

	err := db.C(userCollection).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error Delete User",
			"err":     err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success Delete User",
	})
}
