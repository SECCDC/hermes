package hermes

import (
	"fmt"
	"hermes/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	fmt.Println("Starting Hermes...")
	r := gin.Default()
	r.GET("/healthz", healthCheck)
	r.GET("/api/v1/products", listProducts)

	r.Run()

	defer fmt.Println("Goodbye!")
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func listProducts(c *gin.Context) {
	c.JSON(http.StatusOK, model.Product{
		Name:        "reticulated spline",
		Price:       44.50,
		Description: "It reticulates the splines",
	})
}
