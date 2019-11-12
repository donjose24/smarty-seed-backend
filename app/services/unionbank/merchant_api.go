package unionbank

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmramos02/smarty-seed-backend/app/utils"
	"github.com/jmramos02/smarty-seed-backend/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type PaymentResponse struct {
	Payload   PaymentPayload `json:"payload"`
	Signature string         `json:"signature"`
	Error     string         `json:"error"`
}
type PaymentPayload struct {
	Code        string `json:"code"`
	SenderRefId string `json:"senderRefId"`
	Amount      int    `json:"amount"`
}

func ExecutePayment(amount int, code string) (map[string]interface{}, error) {
	refID := utils.String(15)
	tranDate := time.Now().Format(time.RFC3339Nano)

	body, err := json.Marshal(map[string]interface{}{
		"senderRefId":     refID,
		"tranRequestDate": tranDate[:23],
		"amount": map[string]string{
			"currency": "PHP",
			"value":    strconv.Itoa(amount),
		},
		"remarks":     "Payment Remarks",
		"particulars": "Payment Particulars",
		"info":        []map[string]string{},
	})

	if err != nil {
		panic("Something went wrong: " + err.Error())
	}
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, config.GetUnionbankUrl()+"/merchants/v4/payments/single", bytes.NewBuffer(body))
	request.Header.Add("x-ibm-client-id", config.GetUnionBankClientID())
	request.Header.Add("x-ibm-client-secret", config.GetUnionBankClientSecret())
	request.Header.Add("x-partner-id", config.GetPartnerID())
	request.Header.Add("authorization", "Bearer "+code)
	request.Header.Add("content-type", "application/json")

	resp, err := client.Do(request)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error getting merchant response", err.Error))
	}

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error parsing the response: %v", err.Error))
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading the response: %v", err.Error))
	}

	fmt.Println(response)
	return response, nil
}
