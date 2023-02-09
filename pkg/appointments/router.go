package appointments

import (
	"github.com/egec-org/bookem/pkg/appointments/clients"
	"github.com/gin-gonic/gin"
)

func InitHttpHandlers(r *gin.Engine) {
	r.POST("/appointments", clients.Book)
	r.DELETE("/appointments", clients.Cancel)
	r.GET("/appointments", clients.GetAllAppointments)
	r.GET("/appointments/:id", clients.GetAppointmentById)
	r.GET("/hairstylists/:id/availability_blocks", GetAvailableTimeslots)
	r.GET("hairstylists/:id/prices", ListPrices)
	r.POST("/appointments/confirm", ConfirmAppointment)
}
