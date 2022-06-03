package main

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	addToBalance("DANNON", 500)
	expected := 500
	result := balanceLedger["DANNON"]
	if expected != result {
		t.Errorf("Expected %d, Received %d", result, expected)
	}
}

func TestRemove(t *testing.T) {
	removeFromBalance("MILLER COORS", 100)
	expected := -100
	result := balanceLedger["MILLER COORS"]
	if expected != result {
		t.Errorf("Expected %d, Received %d", result, expected)
	}
}

func TestBalances(t *testing.T) {
	millerCoorsPoints := 500
	dannonPoints := 100
	unileverPoints := 600
	userId := "123"

	balanceLedger["MILLER COORS"] = millerCoorsPoints
	balanceLedger["DANNON"] = dannonPoints
	balanceLedger["UNILEVER"] = unileverPoints
	result := getBalances(userId)
	for _, balance := range result {
		if balance.Payer == "MILLER COORS" {
			if balance.Points != millerCoorsPoints {
				t.Errorf("Expected %d, Received %d", balance.Points, millerCoorsPoints)
			}
		}
		if balance.Payer == "DANNON" {
			if balance.Points != dannonPoints {
				t.Errorf("Expected %d, Received %d", balance.Points, dannonPoints)
			}
		}
		if balance.Payer == "UNILEVER" {
			if balance.Points != unileverPoints {
				t.Errorf("Expected %d, Received %d", balance.Points, unileverPoints)
			}
		}
		if balance.UserId != userId {
			t.Errorf("Expected %q, Received %q", balance.UserId, userId)
		}
	}
}

func TestSpend(t *testing.T) {
	first, _ := time.Parse(time.RFC3339, "2020-10-01T14:00:00Z")
	second, _ := time.Parse(time.RFC3339, "2020-11-01T14:00:00Z")
	third, _ := time.Parse(time.RFC3339, "2020-12-01T14:00:00Z")
	expectedCoorsResult := -100
	expectedDannonResult := -100
	expectedUnileverResult := -300
	transactions = []Transaction{
		{UserId: "123", Payer: "DANNON", Points: 100, Timestamp: first},
		{UserId: "123", Payer: "MILLER COORS", Points: 600, Timestamp: third},
		{UserId: "123", Payer: "UNILEVER", Points: 300, Timestamp: second},
	}
	spendRecords := spendPoints(500)
	for _, spend := range spendRecords {
		if spend.Payer == "MILLER COORS" {
			if spend.Points != expectedCoorsResult {
				t.Errorf("Expected %d, Received %d", spend.Points, expectedCoorsResult)
			}
		}
		if spend.Payer == "DANNON" {
			if spend.Points != expectedDannonResult {
				t.Errorf("Expected %d, Received %d", spend.Points, expectedDannonResult)
			}
		}
		if spend.Payer == "UNILEVER" {
			if spend.Points != expectedUnileverResult {
				t.Errorf("Expected %d, Received %d", spend.Points, expectedUnileverResult)
			}
		}
	}
}
