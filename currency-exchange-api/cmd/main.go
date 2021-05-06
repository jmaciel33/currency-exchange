package main

import (
	"fmt"
	"github.com/jmaciel33/currency-exchange/currency-exchange-api/internal/app/config"
	"github.com/jmaciel33/currency-exchange/currency-exchange-api/internal/app/web"
	"github.com/spf13/viper"
	"log"
)

func main() {
	err:= config.ReadConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	httpPort := fmt.Sprintf(":%s", viper.Get("HTTP_PORT"))

	server := web.Server()
	log.Println(fmt.Sprintf("Service up and running on http://localhost%s", httpPort))
	server.Start(httpPort)
}
