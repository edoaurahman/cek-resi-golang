package expeditions

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func spxTrackingNumber(resi string) string {
	k := "MGViZmZmZTYzZDJhNDgxY2Y1N2ZlN2Q1ZWJkYzlmZDY="
	r := float64(time.Now().UnixNano() / int64(time.Millisecond) / 1e3)
	h := sha256.New()
	rs := fmt.Sprintf("%d", int64(r))
	h.Write([]byte(resi + rs + k))
	return fmt.Sprintf(resi+"|"+rs+"%x", h.Sum(nil))
}

func SpxExpedition(c *gin.Context, resi string) {
	trackingNum := spxTrackingNumber(resi)

	url := fmt.Sprintf("https://spx.co.id/api/v2/fleet_order/tracking/search?sls_tracking_number=%s", trackingNum)
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat permintaan ke API SPX.id"})
		return
	}

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
}
