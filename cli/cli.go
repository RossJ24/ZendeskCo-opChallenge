package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/RossJ24/ZendeskCo-opChallenge/api"
	"github.com/RossJ24/ZendeskCo-opChallenge/models"
)

// Struct representing all the variables the CLI needs to function
type CLI struct {
	cachedTicketList  *[]*models.Ticket
	currentTicketList *[]*models.Ticket
	state             CLIState
	pageNumber        int
	isRunning         bool
	reader            *bufio.Reader
	currentTicket     *models.Ticket
	currentTicketID   int
}

// Default Constructor for instantiating a new CLI struct
func New() CLI {
	return CLI{
		nil,
		nil,
		START,
		1,
		true,
		bufio.NewReader(os.Stdin),
		nil,
		-1,
	}
}

// Updates the state of the CLI
func (cli *CLI) updateState(newState CLIState) {
	cli.state = newState
}

// Prints the error to the user
func (cli CLI) handleError(err error) {
	fmt.Print("The following error occured: ")
	fmt.Println(err.Error())
}

// Run the CLI as a finite state machine
func (cli *CLI) Run() {
	for cli.isRunning {
		switch cli.state {
		case START:
			cli.initialize()
			break
		case SHOW_TICKET_LIST:
			cli.printCurrentList()
			break
		case FETCH_NEXT_TICKET_BATCH:
			cli.getNextTicketBatch()
			break

		case FETCH_PREV_TICKET_BATCH:
			cli.getPreviousTicketBatch()
			break
		case GET_SINGLE_TICKET:
			cli.getTicket()
			break
		case SHOW_TICKET:
			cli.showTicket()
			break
		case MAIN_MENU:
			cli.mainMenu()
			break
		case TICKET_MENU:
			cli.ticketMenu()
			break
		}

	}
	fmt.Println("Thanks for your time! Goodbye.")
}

// Initializes the CLI
func (cli *CLI) initialize() {
	fmt.Println("Welcome!")
	// Load the current ticket list
	jsonObj, err := api.GetTickets(cli.pageNumber, false)
	jsonArray := (*jsonObj)["tickets"].([]interface{})
	// Display errors to the user if there are any
	if err != nil {
		cli.handleError(err)
	}
	cli.currentTicketList = new([]*models.Ticket)
	for _, element := range jsonArray {
		*cli.currentTicketList = append(*cli.currentTicketList, models.TicketFromJSON(element))
	}
	// Update the state to display the newly loaded tickets
	cli.updateState(SHOW_TICKET_LIST)
}
