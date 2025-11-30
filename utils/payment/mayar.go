package payment

import (
	"bytes"
	"encoding/json"
	"go-mayar-payment-webhook/dto"
	"io"
	"net/http"
	"os"
)

func SendMayarInvoice(invoice dto.MayarInvoice) (map[string]interface{}, error) {
	//Crete HTTP request to Mayar API
	apiKey := os.Getenv("MAYAR_API_KEY")
	urlMayar := "https://api.mayar.club/hl/v1/invoice/create"

	// Make header
	headers := map[string]string{
		"Authorization": "Bearer " + apiKey,
		"Content-Type":  "application/json",
	}

	body, err := json.Marshal(invoice)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlMayar, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send HTTP request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Get Response
	bodyResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	// Parsing response to json
	var bodyResponseJSON map[string]interface{}
	if err := json.Unmarshal(bodyResponse, &bodyResponseJSON); err != nil {
		return nil, err
	}

	return bodyResponseJSON, nil
}
