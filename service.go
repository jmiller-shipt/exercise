package main

import "sort"

func addToBalance(payer string, amount int) {
	balanceLedger[payer] = balanceLedger[payer] + amount
}

func removeFromBalance(payer string, amount int) {
	balanceLedger[payer] = balanceLedger[payer] - amount
}

func spendPoints(spendAmount int) []SpendResponse {
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Timestamp.Before(transactions[j].Timestamp)
	})

	m := make(map[string]int)
	var spendReport []SpendResponse
	for _, transaction := range transactions {
		if spendAmount > transaction.Points {
			spendAmount = spendAmount - transaction.Points
			removeFromBalance(transaction.Payer, transaction.Points)
			m[transaction.Payer] = m[transaction.Payer] - transaction.Points
		} else {
			removeFromBalance(transaction.Payer, spendAmount)
			m[transaction.Payer] = m[transaction.Payer] - spendAmount
			break
		}
	}

	for k, v := range m {
		temp := SpendResponse{Payer: k, Points: v}
		spendReport = append(spendReport, temp)
	}

	return spendReport
}

func getBalances() []Balance {
	var balances []Balance

	for k, v := range balanceLedger {
		temp := Balance{Payer: k, Points: v}
		balances = append(balances, temp)
	}
	return balances
}
