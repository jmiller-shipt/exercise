package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/add", add).Methods("POST")
	r.HandleFunc("/spend", spend).Methods("POST")
	r.HandleFunc("/balances", balances).Methods("GET")

	http.ListenAndServe(":80", r)
}

func add(w http.ResponseWriter, r *http.Request) {
	rawBody, err := ioutil.ReadAll(r.Body)

	var transaction Transaction

	if err = json.Unmarshal(rawBody, &transaction); err != nil {
		sendResponse(422, nil, "Unprocessable entity", w)
	} else {
		transactions = append(transactions, transaction)
		addToBalance(transaction.Payer, transaction.Points)
		body := AddResponse{Payer: transaction.Payer}
		sendResponse(201, body, "Successfully added transaction", w)
	}
}

func spend(w http.ResponseWriter, r *http.Request) {
	rawBody, err := ioutil.ReadAll(r.Body)

	var spendRequest SpendRequest

	if err = json.Unmarshal(rawBody, &spendRequest); err != nil {
		sendResponse(422, nil, "Unprocessable entity", w)
	} else {
		spendReport := spendPoints(spendRequest.Points)
		sendResponse(200, spendReport, "Balances after point spend", w)
	}
}

func balances(w http.ResponseWriter, r *http.Request) {
	balances := getBalances()
	sendResponse(200, balances, "Current payer balances", w)
}

func sendResponse(status int, i interface{}, message string, w http.ResponseWriter) {
	response := Response{}
	response.Message = message

	// if error status code
	if status >= 400 {
		response.Success = false
	} else {
		response.Success = true
	}

	response.Data = i

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	marshalledBody, _ := ffjson.Marshal(response)
	w.Write(marshalledBody)
}
