package hermes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	fmt.Println("Starting Hermes...")
	r := gin.Default()
	r.GET("/healthz", healthCheck)
	r.Run()

	defer fmt.Println("Goodbye!")
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
