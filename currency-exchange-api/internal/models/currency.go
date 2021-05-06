package models

type Response struct {
	Code string
	ValueHigh string
}

type Currencies struct {
	USDBRL Currency `json:"USDBRL"`
	EURBRL Currency `json:"EURBRL"`
	INRBRL Currency `json:"INRBRL"`
}

type Currency struct {
	Code string `json:"code"`
	High string `json:"high"`
}
