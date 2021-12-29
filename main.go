package main

import (
	"fmt"
	"strings"
)

var ( // Variables for all three graph symbols
	customerSymbol = " 0 "
	personSymbol   = " * "
	emptySymbol    = " _ "
)

func main() {
	fmt.Println("Welcome to the Target checkout line CLI app. Approach one of our checkout locations")
	fmt.Println("Enter the number of people ahead of you in the line.")
	linePos := 0   // The position in the line for the "customer"
	var people int // Scanning and setting the user input variable, defining the people in front of them in the line
	fmt.Scan(&people)
	people += 1
	graphdimensions := people
	graph := make([]string, graphdimensions) // Making the CLI graphical line representation array
	graph[0] = " 0 "
	for i := 1; i < len(graph); i++ { // Stepping through each of the other line positions and placing a symbol
		graph[i] = personSymbol
	}
	PrintOutGraphic(graph)

	for i := 0; i < len(graph)-1; i++ { // Loop that has a prompt for every open position in front of the customer
		var forwardchar string
		forward := false
		for !forward { //Until the user puts in the correct symbol, it will re-prompt
			fmt.Print("Enter \">\" when you move forward in the line: ")
			fmt.Scan(&forwardchar)
			if forwardchar == ">" { // Checks if the scanned character is correct
				linePos++ //Moves the customer up in the line queue and changes the new symbol and the old symbol
				graph[linePos-1] = emptySymbol
				graph[linePos] = customerSymbol
				forward = true
			}
		}
		PrintOutGraphic(graph)
	}
	fmt.Println("Congradulations! You've reached the front of the line.\n Allow one of our helpful Target representatives help you check out.")
}

func PrintOutGraphic(graphic []string) { // Prints out the graph in string form, with adequate spacing
	fmt.Println()
	fmt.Println(strings.Join(graphic, " "))
	fmt.Println()
}
