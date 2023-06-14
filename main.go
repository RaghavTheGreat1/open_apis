package main

import "fmt"

func main() {
	fmt.Println("1. IFSC to Bank Info")
	fmt.Print("Enter your choice: ")
	var choice int

	choice, err := fmt.Scanf("%d", &choice)

	if err != nil {
		fmt.Print("Please enter a valid choice number.")
		panic(err)
	}

	switch choice {
	case 1:
		ifscToBankInfo()
	}
}
