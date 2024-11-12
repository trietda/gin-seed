package middleware

import (
	"gin-seed/app/auth/service"
	"gin-seed/app/user/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authValues := strings.Split(c.GetHeader("authorization"), "Bearer ")

		if len(authValues) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := authValues[1]
		claim, err := service.VerifyAccessToken(token, &model.UserClaim{})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userClaim", claim)
	}
}
