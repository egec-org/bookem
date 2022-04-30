package main

import (
	"github.com/gin-gonic/gin"
	"github.com/egec-org/bookem/pkg/appointments"
	"github.com/egec-org/bookem/pkg/auth"

)

func main() {
	r := gin.Default()
	r.POST("/appointments", appointments.Book)
	r.DELETE("/appointments", appointments.Cancel)
	r.GET("/appointments", appointments.GetAllAppointments)
	r.GET("/appointments/:id", appointments.GetAppointmentById)
	r.GET("/hairstylists/:id/availability_blocks", appointments.GetAvailableTimeslots)
	r.GET("hairstylists/:id/prices", appointments.ListPrices)
	r.POST("/appointments/confirm", appointments.ConfirmAppointment)
	r.POST("/access/token", auth.AuthUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
