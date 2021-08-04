package cli

import (
	"testing"
)

// Tests that the CLI state is updated correctly when the input is 'next'
func TestMainMenuNext(t *testing.T) {
	cli := New()
	cli.updateState(MAIN_MENU)
	cli.evaluateMainMenuInput("next")
	if cli.state != FETCH_NEXT_TICKET_BATCH {
		t.Errorf("The current state is incorrect.")
	}

}

// Tests that the CLI state is updated correctly when the input is 'prev'
func TestMainMenuPrev(t *testing.T) {
	cli := New()
	cli.updateState(MAIN_MENU)
	cli.evaluateMainMenuInput("prev")
	if cli.state != FETCH_PREV_TICKET_BATCH {
		t.Errorf("The current state is incorrect.")
	}
}

// Tests that the CLI state is updated correctly when the input is a ticket ID
func TestMainMenuID(t *testing.T) {
	cli := New()
	cli.updateState(MAIN_MENU)
	cli.evaluateMainMenuInput("123")
	if cli.state != GET_SINGLE_TICKET {
		t.Errorf("The current state is incorrect.")
	}
}

// Tests that the CLI state is updated correctly when the input is invalid
func TestMainMenuInvalid(t *testing.T) {
	cli := New()
	cli.updateState(MAIN_MENU)
	cli.evaluateMainMenuInput("Invalid")
	if cli.state != MAIN_MENU {
		t.Errorf("The current state is incorrect.")
	}
}
