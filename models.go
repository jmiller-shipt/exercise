package main

import "time"

type Balance struct {
	Payer  string `json:"payer"`
	Points int    `json:"points"`
	UserId string `json:"user_id"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AddResponse struct {
	UserId string `json:"user_id"`
	Payer  string `json:"payer"`
}

type SpendRequest struct {
	UserId string `json:"user_id"`
	Points int    `json:"points"`
}

type SpendRecord struct {
	Payer  string `json:"payer"`
	Points int    `json:"points"`
}

type SpendResponse struct {
	UserId  string        `json:"user_id"`
	Records []SpendRecord `json:"spend_records"`
}

type Transaction struct {
	UserId    string    `json:"user_id"`
	Payer     string    `json:"payer"`
	Points    int       `json:"points"`
	Timestamp time.Time `json:"timestamp"`
}
