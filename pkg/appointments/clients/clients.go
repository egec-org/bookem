package clients

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
)

type Client struct {
	Username string
	Id       int
}

type ClientEntity struct {
	Client
	Password string
}

type ClientRepository interface {
	FetchByUsername(ctx context.Context, username string) (*ClientEntity, error)
}

type ClientService interface {
	FetchByUsername(ctx context.Context, username string) (*Client, error)
}

type ClientHandler interface {
	FetchByUsername() http.Handler
}

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
	conn, err := di.C(c).SafeGet("db-client")
	if err != nil {
		fmt.Println(err)
	}
	// TODO remove fmt statement and fetch appts
	// from DB.
	fmt.Println(conn.(*sqlx.Conn))
	zap.S().Info("Appointments fetched.")
	c.JSON(201, gin.H{
		"message": "Here are your appontments: 1/2 2/4",
	})
}
