package utils

import (
	"testing"

	"github.com/RossJ24/ZendeskCo-opChallenge/models"
)

// Tests that the merge function merges two slices correctly
func TestMerge(t *testing.T) {
	// Crea an expected merged slice
	expected := new([]*models.Ticket)
	for i := 0; i < 6; i++ {
		ticket := &models.Ticket{
			ID: i,
		}
		*expected = append(*expected, ticket)
	}
	// Create a sorted left slice
	left := new([]*models.Ticket)
	for i := 0; i < 3; i++ {
		item := &models.Ticket{
			ID: i,
		}

		*left = append(*left, item)
	}
	// Create a sorted right slice
	right := new([]*models.Ticket)
	for i := 3; i < 6; i++ {
		item := &models.Ticket{
			ID: i,
		}

		*right = append(*right, item)
	}
	output := merge(*left, *right)
	for i := 0; i < 6; i++ {
		if (*expected)[i].ID != output[i].ID {
			t.Errorf("Output: %d, did not equal expected: %d.", output[i].ID, (*expected)[i].ID)
		}
	}
}

// Tests that MergeSort sorts correctly
func TestMergeSort(t *testing.T) {
	// Create sorted expected slice
	expected := new([]*models.Ticket)
	for i := 0; i < 6; i++ {
		ticket := &models.Ticket{
			ID: i,
		}
		*expected = append(*expected, ticket)
	}
	// Create ticket slice with unsorted IDs
	ids := []int{1, 5, 4, 3, 2, 0}
	testTickets := new([]*models.Ticket)
	for i := 0; i < 6; i++ {
		ticket := &models.Ticket{
			ID: ids[i],
		}
		*testTickets = append(*testTickets, ticket)
	}
	output := mergeSort(*testTickets)

	for i := 0; i < 6; i++ {
		if (*expected)[i].ID != output[i].ID {
			t.Errorf("Output: %d, did not equal expected: %d.", output[i].ID, (*expected)[i].ID)
		}
	}
}

// Tests that Binary Search searches correctly
func TestBinarySearch(t *testing.T) {
	// Create Test Current Ticket slice
	testTickets := new([]*models.Ticket)
	for i := 0; i < 6; i++ {
		ticket := &models.Ticket{
			ID: i,
		}
		*testTickets = append(*testTickets, ticket)
	}
	expected := (*testTickets)[1]
	output := binarySearch(*testTickets, 1)
	if (*expected).ID != output.ID {
		t.Errorf("Output: %d, did not equal expected: %d.", output.ID, (*expected).ID)
	}
}

// Tests that the GetLocalTicket will return a non nil ticket pointer if the ticket is local
func TestGetLocalTicketContains(t *testing.T) {
	// Create Test Ticket Cache slice
	testTicketsCache := new([]*models.Ticket)
	for i := 0; i < 6; i++ {
		ticket := &models.Ticket{
			ID: i,
		}
		*testTicketsCache = append(*testTicketsCache, ticket)
	}
	// Create Test Current Ticket slice
	testTickets := new([]*models.Ticket)
	for i := 6; i < 12; i++ {
		ticket := &models.Ticket{
			ID: i,
		}
		*testTickets = append(*testTickets, ticket)
	}
	output := GetLocalTicket(testTickets, testTicketsCache, 4)
	if output == nil {
		t.Errorf("Output pointer was nil.")
	}
	if 4 != output.ID {
		t.Errorf("Output: %d, did not equal expected: 4.", (*output).ID)
	}
}

//Tests that GetLocalTicket will return a nil pointer if the ticket is not local
func TestGetLocalTicketDoesNotContain(t *testing.T) {
	// Create Test Ticket Cache slice
	testTicketsCache := new([]*models.Ticket)
	for i := 0; i < 6; i++ {
		ticket := &models.Ticket{
			ID: i,
		}
		*testTicketsCache = append(*testTicketsCache, ticket)
	}
	// Create Test Current Ticket slice
	testTickets := new([]*models.Ticket)
	for i := 6; i < 12; i++ {
		ticket := &models.Ticket{
			ID: i,
		}
		*testTickets = append(*testTickets, ticket)
	}
	output := GetLocalTicket(testTickets, testTicketsCache, 25)
	if output != nil {
		t.Errorf("Output pointer was not nil.")
	}
}

// Tests the GetLocalTicket will return a nil pointer if there is a nil ticket in the slices that are passed in.
func TestGetLocalTicketWithNull(t *testing.T) {
	// Create Test Ticket Cache slice
	testTicketsCache := new([]*models.Ticket)
	for i := 0; i < 6; i++ {
		*testTicketsCache = append(*testTicketsCache, nil)
	}
	// Create Test Current Ticket slice
	testTickets := new([]*models.Ticket)
	for i := 6; i < 12; i++ {
		ticket := &models.Ticket{
			ID: i,
		}
		*testTickets = append(*testTickets, ticket)
	}
	output := GetLocalTicket(testTickets, testTicketsCache, 25)
	if output != nil {
		t.Errorf("Output pointer was not nil.")
	}
}

// Tests that GetLocalTicket with a Null cache returns proper output
func TestGetLocalTicketWithNullCache(t *testing.T) {
	// Create Test Ticket Cache slice
	testTicketsCache := new([]*models.Ticket)
	testTicketsCache = nil
	// Create Test Current Ticket slice
	testTickets := new([]*models.Ticket)
	for i := 6; i < 12; i++ {
		ticket := &models.Ticket{
			ID: i,
		}
		*testTickets = append(*testTickets, ticket)
	}
	output := GetLocalTicket(testTickets, testTicketsCache, 7)
	expected := (*testTickets)[1]
	if output == nil {
		t.Errorf("Output pointer was nil.")
	}
	if expected.ID != output.ID {
		t.Errorf("Output: %d, did not equal expected: 7.", (*output).ID)
	}

}
