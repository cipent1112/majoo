package router

import (
	"github.com/cipent1112/majoo/auth"
	"github.com/cipent1112/majoo/handler"
	"github.com/gin-gonic/gin"
)

func Route(handler *handler.DB) {
	router := gin.Default()

	router.POST("/login", auth.Login)
	router.GET("/user/:id", auth.Auth, handler.GetUser)
	router.GET("/users", auth.Auth, handler.GetUsers)
	router.POST("/user", auth.Auth, handler.CreateUser)
	router.PUT("/user/:id", auth.Auth, handler.UpdateUser)
	router.DELETE("/user/:id", auth.Auth, handler.DeleteUser)

	router.Run(":3000")
}
