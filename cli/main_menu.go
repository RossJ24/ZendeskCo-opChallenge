package cli

import (
	"fmt"
	"strconv"

	"github.com/RossJ24/ZendeskCo-opChallenge/utils"
)

// Menu that is displayed when the user is viewing a ticket list
func (cli *CLI) mainMenu() {
	// prompt the user for input
	fmt.Println("\nEnter 'next' to view the next page of Ticket results.")
	fmt.Println("Enter 'prev' to view the previous page of Ticket results.")
	fmt.Println("Enter a ticket id to view the Ticket in detail")
	fmt.Println("Enter 'exit' to exit the CLI")
	input, err := cli.reader.ReadString('\n')
	// Inform the user of any error
	if err != nil {
		fmt.Printf("The following error occured while reading your input: %s.\n", err.Error())
		return
	}
	// Clean input
	input = utils.CleanInput(input)
	// Evaluate Input
	cli.evaluateMainMenuInput(input)
}

// Evaluates the input from the user, and updatres the CLI's state accordingly
func (cli *CLI) evaluateMainMenuInput(input string) {
	switch input {
	case "next":
		// Update the state of the CLI to get the next ticket back if the input is 'next'
		cli.updateState(FETCH_NEXT_TICKET_BATCH)
		break
	case "prev":
		// Update the state of the CLI to get the next ticket back if the input is 'prev'
		cli.updateState(FETCH_PREV_TICKET_BATCH)
		break
	case "exit":
		// Update the isRunning variable of the CLI to termninate if the input is 'exit'
		cli.isRunning = false
		break
	// Assume the input is ticket ID if the input is anything else
	default:
		// Parse inout into a int for the ticket ID
		id, err := strconv.Atoi(input)
		// Display any error to the user
		if err != nil {
			fmt.Printf("The following error occured while reading your input: %s.\n", err.Error())
			return
		}
		// If the user inputs an int that is less than 0, consider it invalid input
		if id < 0 {
			fmt.Printf("Ticket IDs must be greater than 0. Your input: %d\n", id)
			return
		}
		// If the id is valid assign it as the current ticket ID
		cli.currentTicketID = id
		// Update the state to fetch the ticket
		cli.updateState(GET_SINGLE_TICKET)
		break
	}
}
