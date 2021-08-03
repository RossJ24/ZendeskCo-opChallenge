package urls

import "fmt"

// Base Zendesk URL
const BASE_URL = "https://zccrossj.zendesk.com"

// URL to get nubmer of Tickets
const COUNT_TICKETS = BASE_URL + "/api/v2/tickets/count.json"

// URL to get Tickets in a list
const GET_TICKETS = BASE_URL + "/api/v2/tickets.json"

//URL tp GET Ticket
const GET_TICKET = BASE_URL + "/api/v2/tickets/"

// Returns the paginated ticket urls with the supplied page number (offset) and the number returned per page (limit)
func GetPaginatedTicketsURL(offset int, limit int) string {
	return fmt.Sprintf("%s?page=%d&per_page=%d", GET_TICKETS, offset, limit)
}

// Returns the url for the given ticket
func GetTicketURL(ticketID int) string {
	return fmt.Sprintf("%s%d.json", GET_TICKET, ticketID)
}
