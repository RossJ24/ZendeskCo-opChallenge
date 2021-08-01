package auth

import (
	"encoding/base64"
	"net/http"
	"os"
)

func BasicAuth(email string, password string) string {
	credentials := email + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials))
}

func AddBasicAuth(req *http.Request) {
	req.Header.Add("Authorization", BasicAuth(os.Getenv("EMAIL"), os.Getenv("PASSWORD")))
}
