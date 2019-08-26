package main

import (
	"reflect"

	"testing"
)

var trackdown *BTCTrackdown

func init() {
	trackdown = NewBTCTrackdown().WithDebug()
}

func TestEmptyTrackdown(t *testing.T) {

	_, err := trackdown.GetTracking("")

	if err == nil {

		t.Error("Transactions without source should fail")
	}
}

func TestTrackdownWithTransactions(t *testing.T) {
	transactions, err := trackdown.GetTracking("3NK9pxgHedLTsrgq1Wkwvx61h9hPt8VrNP")

	if err != nil {
		t.Error("Should not be failed " + err.Error())
	}

	transactionsExpected := []Transaction{
		Transaction{"3Ercgcpbyyrxy7TAVzhoPbeFAGN9tJ8Q4x", 998156},
		Transaction{"3JCVow7NNhAfMFEGLSaEKEUqJPA2AMV2Gf", 12000},
	}

	if !reflect.DeepEqual(transactions, transactionsExpected) {
		t.Errorf("Expected transactions: %v, got %v", transactionsExpected, transactions)
	}
}

func TestTrackdownWithEmptyTransactions(t *testing.T) {
	transactions, err := trackdown.GetTracking("157Bb1npdTvuzth7e4z6mBqzPANKd7QZHB")

	if err != nil {
		t.Error("Should not be failed " + err.Error())
	}

	if (len(transactions)) != 0 {
		t.Error("Transactions with source 157Bb1npdTvuzth7e4z6mBqzPANKd7QZHB should be 0")
	}
}
