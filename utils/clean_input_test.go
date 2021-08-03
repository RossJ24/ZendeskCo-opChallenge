package utils

import (
	"testing"
)

// Tests that correctly formatted input will not be changes
func TestPerfectInput(t *testing.T) {
	expected := "test"
	input := "test"
	output := CleanInput(input)
	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}

}

// Tests that input with a CRLF will become correctly formatted
func TestInputWithCRLF(t *testing.T) {
	expected := "test"
	input := "test\r\n"
	output := CleanInput(input)
	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}
}

// Tests that input with a CRLF and LF will become correctly formatted
func TestInputWithLF(t *testing.T) {
	expected := "test"
	input := "test\n"
	output := CleanInput(input)
	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}
}

// Tests that input with a CRLF, LF, and another LF will become correctly formatted
func TestInputWithCRLFLF(t *testing.T) {
	expected := "test"
	input := "test\r\n\n"
	output := CleanInput(input)
	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}
}

// Tests that input with an LF, CRLF, LF, and another LF will become correctly formatted
func TestInputWithLFCRLFLF(t *testing.T) {
	expected := "test"
	input := "test\n\r\n\n"
	output := CleanInput(input)
	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}
}
