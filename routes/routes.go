package routes

import (
	"example.com/rest_api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEventById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)
	authenticated.POST("/event", createEvent)

	authenticated.POST("/event/:id/register", registerForEvent)
	authenticated.DELETE("/event/:id/register", unregisterForEvent)

	// server.POST("/event", createEvent)
	// server.PUT("/event/:id", updateEvent)
	// server.DELETE("/event/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
