package business

import (
	"encoding/json"
	"fmt"
	"github.com/jmaciel33/currency-exchange/currency-exchange-api/internal/app/config"
	"github.com/jmaciel33/currency-exchange/currency-exchange-api/internal/models"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type CurrencyService struct {
	cacheRepository CacheRepository
}

func NewCurrencyService(repository CacheRepository) CurrencyService {
	return CurrencyService{repository}
}

func (c CurrencyService) GetCurrencies() ([]models.Currency, error) {

	config.ReadConfig()
	var response models.Currencies
	list := make([]models.Currency, 0)
	currenciesFromCache, _ := c.cacheRepository.Get("currency")

	if currenciesFromCache == "" {
		currenciesId := viper.Get("currencies")
		url := "https://economia.awesomeapi.com.br/json/last/"+ fmt.Sprintf("%v", currenciesId)
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Print(err.Error())
		}
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		resp, err := client.Do(req)
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Print(err.Error())
		}

		json.Unmarshal(bodyBytes, &response)

		responseJson, _ := json.Marshal(response)
		c.cacheRepository.Set("currency", string(responseJson), 30)

	} else {
		json.Unmarshal([]byte(currenciesFromCache), &response)
	}

	list =append(list, models.Currency{Code: response.USDBRL.Code, High: response.USDBRL.High})
	list =append(list, models.Currency{Code: response.EURBRL.Code, High: response.EURBRL.High})
	list =append(list, models.Currency{Code: response.INRBRL.Code, High: response.INRBRL.High})

	fmt.Printf("API Response as struct %+v\n", response)

	return list, nil
}
