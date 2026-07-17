package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent)

	authenticate := server.Group("/")
	authenticate.Use(middlewares.Aunthenticate)
	authenticate.POST("/events", createEvent)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)
	authenticate.POST("/events/:id/register")
	authenticate.DELETE("/events/:id/register")

	server.POST("/signup", signup)
	server.POST("/login", login)
}
