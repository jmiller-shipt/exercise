package main

import "time"

type Balance struct {
	Payer  string `json:"payer"`
	Points int    `json:"points"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AddResponse struct {
	Payer string `json:"retailer"`
}

type SpendRequest struct {
	Points int `json:"points"`
}

type SpendResponse struct {
	Payer  string `json:"payer"`
	Points int    `json:"points"`
}

type Transaction struct {
	Payer     string    `json:"payer"`
	Points    int       `json:"points"`
	Timestamp time.Time `json:"timestamp"`
}
