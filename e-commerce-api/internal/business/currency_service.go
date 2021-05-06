package business

import (
	"encoding/json"
	"fmt"
	"github.com/jmaciel33/currency-exchange/e-commerce-api/internal/models"
	"io/ioutil"
	"net/http"
)

func GetCurrencies() ([]models.Currency, error)  {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://currency-exchange-api:8082/v1/currency", nil)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	var response []models.Currency
	json.Unmarshal(bodyBytes, &response)

	fmt.Printf("API Response as struct %+v\n", response)

	return response, nil
}