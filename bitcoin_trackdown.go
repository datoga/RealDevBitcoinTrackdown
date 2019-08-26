package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"github.com/go-resty/resty/v2"
)

type BTCTrackdown struct {
	client *resty.Client
	debug  bool
}

func NewBTCTrackdown() *BTCTrackdown {
	client := resty.New()

	return &BTCTrackdown{
		client: client,
		debug:  false,
	}
}

func (trackdown *BTCTrackdown) WithDebug() *BTCTrackdown {
	trackdown.debug = true
	trackdown.client.SetDebug(true)

	return trackdown
}

func (trackdown BTCTrackdown) GetTracking(source string) ([]Transaction, error) {
	if source == "" {
		return nil, errors.New("Source must not be empty")
	}

	apiAddress := "https://blockchain.info/rawaddr/" + source

	if trackdown.debug {
		fmt.Println("Source to query", apiAddress)
	}

	resp, err := trackdown.client.R().Get(apiAddress)

	if err != nil {
		return nil, err
	}

	respAddr := ApiRespRawAddr{}

	err = json.Unmarshal(resp.Body(), &respAddr)

	if err != nil {
		return nil, err
	}

	if trackdown.debug {
		fmt.Printf("Mapped response: %+v\n", respAddr)
	}

	transactions := trackdown.calculateExternalTransactions(&respAddr, source)

	if trackdown.debug {
		fmt.Printf("Transactions unordered: %+v\n", transactions)
	}

	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Value > transactions[j].Value
	})

	if trackdown.debug {
		fmt.Printf("Transactions ordered: %+v\n", transactions)
	}

	return transactions, nil
}

func (trackdown BTCTrackdown) calculateExternalTransactions(apiRespAddr *ApiRespRawAddr, source string) []Transaction {
	transactions := []Transaction{}

	for _, tx := range apiRespAddr.Transactions {

		if len(tx.Outs) == 0 {
			continue
		}

		if len(tx.Inputs) > 1 {
			continue
		}

		for _, out := range tx.Outs {
			if out.Address == source {
				continue
			}

			transaction := Transaction{
				Address: out.Address,
				Value:   out.Value,
			}

			transactions = append(transactions, transaction)
		}
	}

	return transactions
}
