package appointments

import (
  "github.com/gin-gonic/gin"
  "go.uber.org/zap"
)

func ListPrices(c *gin.Context) {
  zap.S().Info("Prices for hairstylists")
  c.JSON(200, gin.H{
    "message": "20 dollars a haircut. 30 dollars for highlights",
  })
}

func ConfirmAppointment(c *gin.Context) {
  zap.S().Info("Appointment confirmed by stylist.")
  c.JSON(201, gin.H{
    "message": "Appointment confirmed by stylist",
  })
}

func AddAppointmentSlots(c *gin.Context) {
  zap.S().Info("Appointment slots added by stylist.")
  c.JSON(201, gin.H{
    "message": "Appointment slots added by stylist",
  })
}

func RemoveAppontmentSlots(c *gin.Context) {
  zap.S().Info("Appointment slots removed by stylist.")
  c.JSON(201, gin.H{
    "message": "Appointment slots removed by stylist",
  })
}

func GetAvailableTimeslots(c *gin.Context) {
  zap.S().Info("Stylist B available for 2/7.")
  c.JSON(201, gin.H{
    "message": "Stylist B available for 2/7.",
  })
}
