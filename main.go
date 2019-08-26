package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	trackdown := NewBTCTrackdown()

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		scanner.Scan()

		source := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from input: ", err)
			continue
		}

		transactions, err := trackdown.GetTracking(source)

		if err != nil {
			fmt.Println("Error getting the source: ", err)
			continue
		}

		response := (TransactionResponse)(transactions)

		fmt.Println(response)
	}
}
