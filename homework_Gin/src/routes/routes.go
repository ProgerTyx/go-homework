package routes

import (
	"github.com/gin-gonic/gin"
	handle "go-homework/homework_Gin/src/handlers"
	middleware "go-homework/homework_Gin/src/middlewares"
	services "go-homework/homework_Gin/src/services"
	"net/http"
)

type Routes struct{}

func (c Routes) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/users", middleware.AuthorizeJWT(), handle.GetAllUser)
		api.GET("/users/:id", middleware.AuthorizeJWT(), handle.GetUser)
		api.POST("/login", func(ctx *gin.Context) {
			token := services.Login(ctx)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}
		})

		api.POST("/users", middleware.AuthorizeJWT(), handle.CreateUser)
		api.PUT("/users/:id", middleware.AuthorizeJWT(), handle.UpdateUser)
		api.DELETE("/users/:id", middleware.AuthorizeJWT(), handle.DeleteUser)

		//api.GET("/create-admin", handle.CreateAdmin)
	}
	r.Run(":8000")
}
