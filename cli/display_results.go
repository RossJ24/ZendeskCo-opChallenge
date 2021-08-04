package cli

import "fmt"

// Printss the current list of tickets
func (cli *CLI) printCurrentList() {
	// Print the header for a ticket page
	fmt.Println("Tickets:")
	fmt.Println("--------------------")
	for _, element := range *cli.currentTicketList {
		// Print the shortened representation of the ticket if it is not nil
		if element != nil {
			fmt.Println(element.GetListInfo())
		} else {
			// If the ticket is nil, inform the user that there was an error
			fmt.Println("Unable to Parse Ticket JSON")
		}
	}
	cli.updateState(MAIN_MENU)
}

// Prints the current Ticket
func (cli *CLI) showTicket() {
	// Get detailed string representation of the Ticket
	ticket := cli.currentTicket.GetIndividualDisplayInfo()
	// Print the ticket
	fmt.Println()
	fmt.Println(ticket)
	fmt.Println()
	// Update the CLI's state to display the ticket menu
	cli.updateState(TICKET_MENU)
}
