package main

import (
	"bytes"
	"fmt"
)

type Transaction struct {
	Address string
	Value   int
}

type TransactionResponse []Transaction

func (response TransactionResponse) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("[")

	for i, tx := range response {
		s := fmt.Sprintf("[\"%s\", %d]", tx.Address, tx.Value)

		buffer.WriteString(s)

		if i != len(response)-1 {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("]")

	return buffer.String()
}
