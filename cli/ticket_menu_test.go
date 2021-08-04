package cli

import "testing"

// Tests that the CLI state is updated correctly when the input is 'back'
func TestTicketMenuBack(t *testing.T) {
	cli := New()
	cli.updateState(TICKET_MENU)
	cli.evaluateTicketMenuInput("back")
	if cli.state != SHOW_TICKET_LIST {
		t.Errorf("The current state is incorrect.")
	}
}

// Tests that the CLI state is updated correctly when the input is 'exit'
func TestTicketMenuExit(t *testing.T) {
	cli := New()
	cli.updateState(TICKET_MENU)
	cli.evaluateTicketMenuInput("exit")
	if cli.isRunning {
		t.Errorf("CLI is still running")
	}
}

// Tests that the CLI state is updated correctly when the input is invalid
func TestTicketMenuInvalid(t *testing.T) {
	cli := New()
	cli.updateState(TICKET_MENU)
	cli.evaluateTicketMenuInput("invalid")
	if cli.state != TICKET_MENU {
		t.Errorf("The current state is incorrect.")
	}
}
