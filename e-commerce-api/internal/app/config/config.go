package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func ReadConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	return nil
}

func InitDbConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connectionString())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectionString() string {

	ReadConfig()

	host := viper.Get("CRDB.CRDB_HOST")
	port := viper.Get("CRDB.CRDB_PORT")
	crdb := viper.Get("CRDB.CRDB_DATABASE")
	user := viper.Get("CRDB.CRDB_USER")
	pass := viper.Get("CRDB.CRDB_PASS")

	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, crdb, user, pass)
}

func InitSchemaDatabase() {

	db, err := InitDbConnection()
	if err != nil {
		panic(err)
	}

	db.MustExec("CREATE DATABASE IF NOT EXISTS sbfcase_products")
	db.MustExec("CREATE TABLE IF NOT EXISTS sbfcase_products.products (product_id UUID PRIMARY KEY, name String NOT NULL, price DECIMAL NOT NULL)")

}