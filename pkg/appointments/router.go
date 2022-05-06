package appointments

import (
  "github.com/gin-gonic/gin"
)

func InitHttpHandlers(r *gin.Engine) {
  r.POST("/appointments", Book)
  r.DELETE("/appointments", Cancel)
  r.GET("/appointments", GetAllAppointments)
  r.GET("/appointments/:id", GetAppointmentById)
  r.GET("/hairstylists/:id/availability_blocks", GetAvailableTimeslots)
  r.GET("hairstylists/:id/prices", ListPrices)
  r.POST("/appointments/confirm", ConfirmAppointment)
}
