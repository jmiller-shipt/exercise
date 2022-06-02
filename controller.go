package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pquerna/ffjson/ffjson"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/add", add).Methods("POST")
	r.HandleFunc("/spend", spend).Methods("POST")
	r.HandleFunc("/balances/{userId}", balances).Methods("GET")

	err := http.ListenAndServe(":80", r)
	if err != nil {
		fmt.Print("Could not start web server")
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	rawBody, err := ioutil.ReadAll(r.Body)

	var transaction Transaction

	if err = json.Unmarshal(rawBody, &transaction); err != nil {
		sendResponse(422, nil, "Unprocessable entity", w)
	} else {
		transactions = append(transactions, transaction)
		addToBalance(transaction.Payer, transaction.Points)
		body := AddResponse{Payer: transaction.Payer, UserId: transaction.UserId}
		sendResponse(201, body, "Successfully added transaction", w)
	}
}

func spend(w http.ResponseWriter, r *http.Request) {
	rawBody, err := ioutil.ReadAll(r.Body)

	var spendRequest SpendRequest

	if err = json.Unmarshal(rawBody, &spendRequest); err != nil {
		sendResponse(422, nil, "Unprocessable entity", w)
	} else {
		spendResponse := spendPoints(spendRequest.Points)
		sendResponse(200, spendResponse, "Points spent by payer", w)
	}
}

func balances(w http.ResponseWriter, r *http.Request) {
	userId, _ := mux.Vars(r)["userId"]
	balances := getBalances(userId)
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
