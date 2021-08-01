package auth

import (
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	godotenv.Load("../../.env")
	os.Exit(m.Run())
}

// Tests the encoding for basic auth (Pretty much Base64)
func TestBasicAuth(t *testing.T) {
	output := BasicAuth("email", "password")
	expected := "Basic " + "ZW1haWw6cGFzc3dvcmQ="

	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}

}

// Tests that the basic auth string is being put into the correct header
func TestAddBasicAuth(t *testing.T) {
	req, _ := http.NewRequest("GET", "URL", nil)
	AddBasicAuth(req)
	expected := BasicAuth(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	output := req.Header.Get("Authorization")
	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}
}
