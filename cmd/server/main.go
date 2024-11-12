package main

import (
	"gin-seed/app/auth/middleware"
	userhandler "gin-seed/app/user/handler"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	api := r.Group("api")

	user := api.Group("users")
	user.POST("/sessions", userhandler.Login)
	user.POST("/", userhandler.Register)
	user.GET("/me", middleware.JwtAuth(), userhandler.Test)

	r.Run(":3000")
}
