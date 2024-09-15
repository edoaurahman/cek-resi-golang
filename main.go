package main

import (
	"fmt"
	"net/http"
	"os"
	"spx-tracker/expeditions"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	inisialization()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	port := os.Getenv("PORT")
	fmt.Println("Server is running on port", port)
	r.GET("/cekresi", func(c *gin.Context) {
		resi := c.Query("sls_tracking_number")
		expedition := c.Query("type")
		fmt.Println(expedition)
		if resi == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'sls_tracking_number' harus diisi"})
			return
		}

		switch expedition {
		case "spx":
			expeditions.SpxExpedition(c, resi)
		case "jnt-cargo":
			expeditions.JntCargoExpedition(c, resi)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Expedisi tidak dikenal"})
		}

	})

	r.Run(":" + port)
}

// initialization function
func inisialization() {
	LoadEnv()
}

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
