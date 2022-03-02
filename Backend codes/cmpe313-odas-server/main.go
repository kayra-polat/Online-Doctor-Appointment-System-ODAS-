package main

import (
	"cmpe313-odas-server/controllers"
	"cmpe313-odas-server/middlewares"
	"cmpe313-odas-server/models"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
	"net/http"
)

func initHandler() *gin.Engine {
	r := gin.Default()
	r.Use(cors.AllowAll())

	v1 := r.Group("/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	appointments := v1.Group("/appointments")
	{
		appointments.GET("/", middlewares.ValidateAuthorization("Both"), controllers.ReadAllAppointments)
		appointments.GET("/:id", middlewares.ValidateAuthorization("Both"), controllers.ReadAppointment)
		appointments.POST("/", middlewares.ValidateAuthorization("Patient"), controllers.CreateAppointment)
		appointments.DELETE("/:id", middlewares.ValidateAuthorization("Patient"), controllers.DeleteAppointment)
	}
	availables := v1.Group("/availables")
	{
		availables.GET("/", middlewares.ValidateAuthorization("Both"), controllers.ReadAllAvailableTime)
		availables.POST("/", middlewares.ValidateAuthorization("Doctor"), controllers.CreateAvailableTime)
		availables.DELETE("/:id", middlewares.ValidateAuthorization("Doctor"), controllers.DeleteAvailableTime)
	}
	profiles := v1.Group("/profiles", middlewares.ValidateAuthorization("Both"))
	{
		profiles.GET("/:id", controllers.ReadProfile)
	}

	return r
}

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: initHandler(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Cannot start server: Reason: %s", err.Error())
	}

	_ = models.GetDatabaseInstance()
}
