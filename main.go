package main

import (
	"fmt"
	"net/http"
	"spx-tracker/expeditions"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

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

	r.Run(":3000") // Ganti dengan port yang sesuai dengan kebutuhan Anda
}
