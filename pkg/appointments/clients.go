package appointments

import (
  "github.com/gin-gonic/gin"
  "go.uber.org/zap"
)

func Book(c *gin.Context) {
  zap.S().Info("Booked and saved appointment in DB")
  c.JSON(201, gin.H{
    "message": "Your appointment is booked.",
  })
}

func Cancel(c *gin.Context) {
  zap.S().Info("Appointment canceled.")
  c.JSON(201, gin.H{
    "message": "Your appointment has been canceled.",
  })
}

func GetAppointmentById(c *gin.Context) {
  zap.S().Info("Found appointment 123")
  c.JSON(201, gin.H{
    "message": "Found appointment 123.",
  })
}

func GetAllAppointments(c *gin.Context) {
  zap.S().Info("Appointments fetched.")
  c.JSON(201, gin.H{
    "message": "Here are your appontments: 1/2 2/4",
  })
}
