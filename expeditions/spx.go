package expeditions

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"spx-tracker/models"
	"time"

	"github.com/gin-gonic/gin"
)

func spxTrackingNumber(resi string) string {
	k := os.Getenv("SPX_TOKEN")
	r := float64(time.Now().UnixNano() / int64(time.Millisecond) / 1e3)
	h := sha256.New()
	rs := fmt.Sprintf("%d", int64(r))
	h.Write([]byte(resi + rs + k))
	return fmt.Sprintf(resi+"|"+rs+"%x", h.Sum(nil))
}

// Spx expeditions
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

	responseSpx(c, resp, responseData)
}

func responseSpx(c *gin.Context, res *http.Response, body []byte) {
	var response models.Response
	var model models.SpxModel

	err := json.Unmarshal([]byte(body), &model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data respons dari API SPX.id"})
		return
	}

	response.Resi = model.Data.SlsTrackingNumber
	response.Expedition = "SPX"
	for _, detail := range model.Data.TrackingList {
		response.Details = append(response.Details, models.Details{
			Time:    time.Unix(int64(detail.Timestamp), 0),
			Message: detail.Message,
		})
	}

	c.JSON(res.StatusCode, response)
}
