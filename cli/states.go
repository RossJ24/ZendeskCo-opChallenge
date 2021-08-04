package cli

type CLIState int

// States for the CLI
const (
	START CLIState = iota
	FETCH_PREV_TICKET_BATCH
	FETCH_NEXT_TICKET_BATCH
	SHOW_TICKET
	SHOW_TICKET_LIST
	GET_SINGLE_TICKET
	MAIN_MENU
	TICKET_MENU
)
