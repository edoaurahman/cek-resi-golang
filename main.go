package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func trackingNumber(resi string) string {
	k := "MGViZmZmZTYzZDJhNDgxY2Y1N2ZlN2Q1ZWJkYzlmZDY="
	r := math.Floor(float64(time.Now().UnixNano() / int64(time.Millisecond) / 1e3))
	h := sha256.New()
	rs := fmt.Sprintf("%d", int64(r))
	h.Write([]byte(resi + rs + k))
	return fmt.Sprintf(resi+"|"+rs+"%x", h.Sum(nil))
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/cekresi", func(c *gin.Context) {
		resi := c.Query("sls_tracking_number")
		if resi == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'sls_tracking_number' harus diisi"})
			return
		}

		trackingNum := trackingNumber(resi)

		// Lakukan request ke API SPX.id dengan menggunakan trackingNum sebagai parameter
		// Disini Anda harus mengganti {API_KEY_ANDA} dengan API key yang Anda miliki dari SPX.id
		url := fmt.Sprintf("https://spx.co.id/api/v2/fleet_order/tracking/search?sls_tracking_number=%s", trackingNum)
		client := http.DefaultClient
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat permintaan ke API SPX.id"})
			return
		}

		req.Header.Set("Authorization", "{API_KEY_ANDA}") // Ganti {API_KEY_ANDA} dengan API key Anda
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan respons dari API SPX.id"})
			return
		}
		defer resp.Body.Close()

		// Baca data dari respons API dan kirimkan sebagai respons ke klien (client)
		var responseData []byte
		responseData, err = io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data respons dari API SPX.id"})
			return
		}

		c.Data(resp.StatusCode, "application/json", responseData)
	})

	r.Run(":3000") // Ganti dengan port yang sesuai dengan kebutuhan Anda
}
