package urls

import (
	"testing"
)

// Tests that the paginated ticket URL is being constructed correctly
func TestGetPaginatedTicketsURL(t *testing.T) {
	output := GetPaginatedTicketsURL(1, 25)
	expected := "https://zccrossj.zendesk.com/api/v2/tickets.json?page=1&per_page=25"
	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}
}

// Tests that the ticket URL is being constructed correctly
func TestGetTicketURL(t *testing.T) {
	output := GetTicketURL(1)
	expected := "https://zccrossj.zendesk.com/api/v2/tickets/1.json"
	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}
}
