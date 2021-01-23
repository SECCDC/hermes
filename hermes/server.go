package hermes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hermes/model"
)

type Config struct {
	DBUser string
	DBPass string
	DBAddr string
	DBName string
}

func Run(c Config) {
	fmt.Println("Starting Hermes...")

	dbConnect(c.DBUser, c.DBPass, c.DBAddr, c.DBName)

	r := gin.Default()
	r.GET("/healthz", healthCheck)
	r.GET("/api/v1/products", listProducts)
	r.Run()

	defer fmt.Println("Goodbye!")
}

func dbConnect(user, pass, address, dbName string) *gorm.DB {
	fmt.Println("I got an address: " + address)
	fmt.Println("I got an password: " + pass)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, address, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
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
