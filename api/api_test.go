package api

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	godotenv.Load("../.env")
	os.Exit(m.Run())
}

func TestGetTicketsLoud(t *testing.T) {
	// Original Strerr file descriptor (To Reassign it later)
	original := os.Stderr
	// Replace Stderr witt the file descriptors of the pipe
	r, w, _ := os.Pipe()
	os.Stderr = w
	// Create a channel to write the Stderr output to
	outChan := make(chan string)
	// Create a goroutine before printing so that printing does not block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outChan <- buf.String()
	}()
	// Test that tickets are fetched successfully
	_, err := GetTickets(1, false)
	if err != nil {
		t.Errorf("Unable to fetch tickets. Error: %s", err.Error())
	}
	// Close the write file descriptor
	w.Close()
	// Restore Stderr to the original file descriptor
	os.Stderr = original
	// Read the contents of the output channel
	out := <-outChan
	// Check the length of the output
	if len(out) == 0 {
		t.Errorf("Output is empty. No progress bar was rendered.")
	}
}

func TestGetTicketsSilent(t *testing.T) {
	// Original Strerr file descriptor (To Reassign it later)
	original := os.Stderr
	// Replace Stderr witt the file descriptors of the pipe
	r, w, _ := os.Pipe()
	os.Stderr = w
	// Create a channel to write the Stderr output to
	outChan := make(chan string)
	// Create a goroutine before printing so that printing does not block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outChan <- buf.String()
	}()
	// Test that tickets are fetched successfully
	_, err := GetTickets(1, true)
	if err != nil {
		t.Errorf("Unable to fetch tickets. Error: %s", err.Error())
	}
	// Close the write file descriptor
	w.Close()
	// Restore Stderr to the original file descriptor
	os.Stderr = original
	// Read the contents of the output channel
	out := <-outChan
	if len(out) != 0 {
		t.Errorf("Output is not empty. It has length %d, and contents: %s.", len(out), out)
	}
}

// Tests that the ticket count is being returned correctly
func TestGetTicketCount(t *testing.T) {
	_, err := GetAccountTicketCount()
	if err != nil {
		t.Errorf("Unable to fetch ticket count. Error: %s", err.Error())
	}
}
