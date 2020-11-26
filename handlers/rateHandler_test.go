package handlers

import (
	"testing"
	"time"
)

func TestValidationShouldPass(t *testing.T) {
	tradeDate := time.Now().AddDate(0, 0, +1)
	result := checkIfTradeIsWithinNextThreeDays(tradeDate)

	if result != true {
		t.Error("test failed.")
	}
}

func TestValidationShouldFailWhenTradeDateInFuture(t *testing.T) {
	tradeDate := time.Now().AddDate(0, 0, -1)
	result := checkIfTradeIsWithinNextThreeDays(tradeDate)

	if result != false {
		t.Error("test should failed as trade not within criteria")
	}
}

func TestValidationShouldFailWhenTradeInPast(t *testing.T) {
	tradeDate := time.Now().AddDate(0, 0, +4)
	result := checkIfTradeIsWithinNextThreeDays(tradeDate)

	if result != false {
		t.Error("test should failed as trade not within criteria")
	}
}

func TestHasValidDataShouldReturnTrue(t *testing.T) {
	postData := PostData{
		SourceCurrency:      "USD",
		DestinationCurrency: "GBP",
		Amount:              12.99,
		TradeDate:           time.Now().AddDate(0, 0, 1),
	}

	res := hasValidData(postData)

	if res != true {
		t.Error("test failed")
	}
}

func TestHasValidDataShouldReturnFalse(t *testing.T) {
	postData := PostData{
		SourceCurrency:      "USD",
		DestinationCurrency: "GBP",
		Amount:              12.99,
		TradeDate:           time.Now().AddDate(0, 0, 4),
	}

	res := hasValidData(postData)

	if res != false {
		t.Error("test failed")
	}
}
