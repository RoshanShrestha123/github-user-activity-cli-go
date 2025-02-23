package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the Username of the target github:")
	var username string
	fmt.Scan(&username)

	fmt.Printf("**Searching result for the %s on github**\n", username)

}
