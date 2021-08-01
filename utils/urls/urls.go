package urls

import "fmt"

// Base Zendesk URL
const BASE_URL = "https://zccrossj.zendesk.com"

// URL to get nubmer of Tickets
const COUNT_TICKETS = BASE_URL + "/api/v2/tickets/count.json"

// URL to get Tickets in a list
const GET_TICKETS = BASE_URL + "/api/v2/tickets.json"

// Returns the paginate ticket urls with the supplied page number (offset) and the number returned per page (limit)
func GetPaginatedTicketsURL(offset int32, limit int32) string {
	return fmt.Sprintf("%s?page=%d&per_page=%d", GET_TICKETS, offset, limit)
}
