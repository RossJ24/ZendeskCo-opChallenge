package cli

import (
	"fmt"

	"github.com/RossJ24/ZendeskCo-opChallenge/api"
	"github.com/RossJ24/ZendeskCo-opChallenge/models"
	"github.com/RossJ24/ZendeskCo-opChallenge/utils"
)

// Fetches the next batch of tickets and updates the CLI's variables
func (cli *CLI) getNextTicketBatch() {
	// If the last ticket list wasn't full, stop the user from going forward
	if len(*cli.currentTicketList) < 25 {
		fmt.Println("Cannot go forward")
		cli.updateState(SHOW_TICKET_LIST)
		return
	}
	// Fetch the next page of tickets
	cli.pageNumber++
	jsonObj, err := api.GetTickets(cli.pageNumber, false)
	jsonArray := (*jsonObj)["tickets"].([]interface{})
	if err != nil {
		cli.handleError(err)
		return
	}
	// Update the cached ticket list
	cli.cachedTicketList = cli.currentTicketList
	// Parse the tickets into ticket structs and fill in a new Ticket slice
	cli.currentTicketList = new([]*models.Ticket)
	for _, element := range jsonArray {
		*cli.currentTicketList = append(*cli.currentTicketList, models.TicketFromJSON(element))
	}
	// Update the state to show the newly fetched ticket list
	cli.updateState(SHOW_TICKET_LIST)
}

// Fetches the previous batch of tickets and updates the CLI's variables
func (cli *CLI) getPreviousTicketBatch() {
	// If the CLI is at page one, stop the user from going backward
	if cli.pageNumber == 1 {
		fmt.Println("\nCannot go back")
		cli.updateState(SHOW_TICKET_LIST)
		return
	}
	// decrement the page number
	cli.pageNumber--
	// Update the current list of tickets to the caches list of tickets
	cli.currentTicketList = cli.cachedTicketList
	// Fetch the previous page's previous page of tickets if possible
	if cli.pageNumber > 1 {
		jsonObj, err := api.GetTickets(cli.pageNumber-1, true)
		if err != nil {
			cli.handleError(err)
			cli.updateState(MAIN_MENU)
			return
		}
		jsonArray := (*jsonObj)["tickets"].([]interface{})
		// Parse the tickets into ticket structs and fill in a new Ticket slice
		cli.cachedTicketList = new([]*models.Ticket)
		for _, element := range jsonArray {
			*cli.cachedTicketList = append(*cli.cachedTicketList, models.TicketFromJSON(element))
		}
	} else {
		// If it is not possible to fetch the previous page's previous page of tickets, make it nil
		cli.cachedTicketList = nil
	}
	// Update the state to show the new current ticket batch
	cli.updateState(SHOW_TICKET_LIST)
}

// Get a single ticket as the current Ticket
func (cli *CLI) getTicket() {
	// Search the cache and current ticekt lists for the ticket
	cli.currentTicket = utils.GetLocalTicket(cli.currentTicketList, cli.cachedTicketList, cli.currentTicketID)
	// If the ticket is not in any of the in memory lists fetch from the API
	if cli.currentTicket == nil {
		jsonObj, err := api.GetTicket(cli.currentTicketID)
		// If there are any errors present them to the user
		if err != nil {
			cli.handleError(err)
			cli.updateState(MAIN_MENU)
			return
		}

		jsonTicket := (*jsonObj)["ticket"].(map[string]interface{})
		cli.currentTicket = models.TicketFromJSON(jsonTicket)
	}
	// Update the state to shwo the ticket
	cli.updateState(SHOW_TICKET)
}
