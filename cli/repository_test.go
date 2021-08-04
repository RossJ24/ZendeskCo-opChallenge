package cli

import (
	"testing"

	"github.com/RossJ24/ZendeskCo-opChallenge/models"
)

// The tests in this file are e2e tests, not unit tests

// Tests that the CLI updates the current batch on 'next' when it is full
func TestGetNextTicketBatchFull(t *testing.T) {
	cli := New()
	// Simulate the currentTicketList being full
	cli.currentTicketList = new([]*models.Ticket)
	oldPointer := cli.currentTicketList
	for i := 0; i < 25; i++ {
		(*cli.currentTicketList) = append((*cli.currentTicketList), &models.Ticket{ID: i})
	}
	cli.getNextTicketBatch()
	if cli.currentTicketList == oldPointer {
		t.Error("The current ticket list was not updated.")
	}
}

// Tests that the CLI doesn't update the current batch on 'next' when it is not full
func TestGetNextTicketBatchNotFull(t *testing.T) {
	cli := New()
	// Simulate the curent Ticket list not being full
	cli.currentTicketList = new([]*models.Ticket)
	oldPointer := cli.currentTicketList
	for i := 0; i < 20; i++ {
		(*cli.currentTicketList) = append((*cli.currentTicketList), &models.Ticket{ID: i})
	}
	cli.getNextTicketBatch()
	if cli.currentTicketList != oldPointer {
		t.Error("The current ticket list was erroneously updated.")
	}
}

// Tests that the CLI updates the current batch correctly on 'prev' when the cache is not nilf
func TestGetPreviousTicketBatchNotNil(t *testing.T) {
	cli := New()
	// Simulate the CLI being on page 2
	cli.pageNumber++
	cli.cachedTicketList = new([]*models.Ticket)
	for i := 0; i < 25; i++ {
		(*cli.cachedTicketList) = append((*cli.cachedTicketList), &models.Ticket{ID: i})
	}
	cli.currentTicketList = new([]*models.Ticket)
	oldPointer := cli.cachedTicketList
	for i := 0; i < 20; i++ {
		(*cli.currentTicketList) = append((*cli.currentTicketList), &models.Ticket{ID: i})
	}

	cli.getPreviousTicketBatch()
	if cli.currentTicketList != oldPointer {
		t.Error("The current ticket list was erroneously updated.")
	}
}

// Tests that the CLI does not update the current batch on 'prev' when the cache is nil
func TestGetPreviousTicketBatchNil(t *testing.T) {
	cli := New()
	// Simulate the CLI being on page 1
	cli.cachedTicketList = new([]*models.Ticket)
	cli.cachedTicketList = nil
	cli.currentTicketList = new([]*models.Ticket)
	oldPointer := cli.currentTicketList
	for i := 0; i < 20; i++ {
		(*cli.currentTicketList) = append((*cli.currentTicketList), &models.Ticket{ID: i})
	}
	cli.getPreviousTicketBatch()
	if cli.currentTicketList != oldPointer {
		t.Error("The current ticket list was erroneously updated.")
	}
}
