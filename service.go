package main

import "sort"

func addToBalance(payer string, amount int) {
	balanceLedger[payer] = balanceLedger[payer] + amount
}

func removeFromBalance(payer string, amount int) {
	balanceLedger[payer] = balanceLedger[payer] - amount
}

func spendPoints(spendAmount int) []SpendRecord {
	// Sort transactions from oldest to newest
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Timestamp.Before(transactions[j].Timestamp)
	})

	// Initialize map to record spend amounts
	amountsSpent := make(map[string]int)

	var spendReport []SpendRecord

	// Iterate through transactions
	for _, transaction := range transactions {

		// if remaining spend amount is less
		// than current transactions points
		if spendAmount > transaction.Points {
			spendAmount = spendAmount - transaction.Points
			removeFromBalance(transaction.Payer, transaction.Points)
			amountsSpent[transaction.Payer] = amountsSpent[transaction.Payer] - transaction.Points
		} else { // remove only the remaining spend amount
			removeFromBalance(transaction.Payer, spendAmount)
			amountsSpent[transaction.Payer] = amountsSpent[transaction.Payer] - spendAmount
			break
		}
	}

	// Condense amounts spent to payer
	for k, v := range amountsSpent {
		temp := SpendRecord{Payer: k, Points: v}
		spendReport = append(spendReport, temp)
	}

	return spendReport
}

func getBalances(userId string) []Balance {
	var balances []Balance

	// Convert map to Balance structs
	for payer, points := range balanceLedger {
		temp := Balance{Payer: payer, Points: points, UserId: userId}
		balances = append(balances, temp)
	}

	return balances
}
