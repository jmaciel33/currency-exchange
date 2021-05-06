package factory

import (
	"github.com/jmaciel33/currency-exchange/currency-exchange-api/internal/app/config"
	"github.com/jmaciel33/currency-exchange/currency-exchange-api/internal/app/persistence"
	"github.com/jmaciel33/currency-exchange/currency-exchange-api/internal/business"
)

var repository business.CacheRepository

func init()  {
	config.ReadConfig()
	redis := config.RedisClient()
	repository = persistence.NewCacheRepository(redis)
}

func GetCacheService() business.CurrencyService {
	return business.NewCurrencyService(repository)
}