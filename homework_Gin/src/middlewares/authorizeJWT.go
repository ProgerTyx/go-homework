package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-homework/homework_Gin/src/services"
	"net/http"
	"strings"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		tokenString = strings.Trim(tokenString, " ")
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
