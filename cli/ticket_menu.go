package cli

import (
	"fmt"

	"github.com/RossJ24/ZendeskCo-opChallenge/utils"
)

// Menu that is displayed when the user is currently viewing a ticket.
func (cli *CLI) ticketMenu() {
	// Promp the user for input
	fmt.Println("Enter 'back' to go back to the main menu.")
	fmt.Println("Enter 'exit' to exit the CLI")
	input, err := cli.reader.ReadString('\n')
	// Present any error to the user
	if err != nil {
		fmt.Printf("The following error occured while reading your input: %s.\n", err.Error())
	}
	// Clean input
	input = utils.CleanInput(input)
	// Evaluate input
	cli.evaluateTicketMenuInput(input)
}

// Evaluates the input from the user, and updatres the CLI's state accordingly
func (cli *CLI) evaluateTicketMenuInput(input string) {
	switch input {
	// If the input is back, update the state to show the current ticket list
	case "back":
		cli.updateState(SHOW_TICKET_LIST)
		cli.currentTicketID = -1
		cli.currentTicket = nil
		break
	// If the input is exit, set the isRunning variable to false to stop execution
	case "exit":
		cli.isRunning = false
		break
	// Inform the user that their input was invalid if it doesn't match any of the other cases
	default:
		fmt.Printf("Invalid input")
	}
}
