package main

import (
	"fmt"
	"strings"
)

func getUsername() string {
	for {
		var username string
		fmt.Println("Please Enter you username: ")
		_, err := fmt.Scan(&username)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			fmt.Println("Please try again")
			continue
		}

		username = strings.TrimSpace(username)
		if username == "" || len(username) < 4 || len(username) > 10 {
			fmt.Println("Username should be between 4 to 10 characters")
			continue
		}

		return username
	}
}
