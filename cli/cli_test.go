package cli

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	godotenv.Load("../.env")
	os.Exit(m.Run())
}

// Tests that the CLI initializes properly
func TestInitialize(t *testing.T) {
	cli := New()
	cli.initialize()
	if cli.currentTicketList == nil {
		t.Error("currentTicketList is nil.")
	}
	if cli.state == MAIN_MENU {
		t.Error("Incorrect state following initialization")
	}
}
