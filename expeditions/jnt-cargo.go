package expeditions

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JntCargoExpedition(c *gin.Context, resi string) {
	url := "https://office.jtcargo.co.id/official/waybill/trackingCustomerByWaybillNo"
	method := "POST"

	payload := strings.NewReader(`{
    "waybillNo": "` + resi + `",
    "langType": "ID",
    "searchWaybillOrCustomerOrderId": "1"
}`)

	client := http.DefaultClient
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan respons dari API J&T Cargo"})
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data respons dari API J&T Cargo"})
		return
	}
	c.Data(res.StatusCode, "application/json", body)
}
