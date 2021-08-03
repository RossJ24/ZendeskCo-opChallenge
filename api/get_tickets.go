package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/RossJ24/ZendeskCo-opChallenge/utils/auth"
	"github.com/RossJ24/ZendeskCo-opChallenge/utils/urls"
	"github.com/schollz/progressbar/v3"
)

// The tests in this file are not unit tests, but end-to-end (e2e) tests

// Limit of the number of tickets that can be fetched in a single call to the tickets API
const LIMIT = 25

// Gets Tickets in JSON string format and parses it into an arbitrary map
func GetTickets(offset int, silent bool) (*map[string]interface{}, error) {
	// Get paginated URL to fetch tickets
	url := urls.GetPaginatedTicketsURL(offset, 25)
	// Create request struct
	req, err := http.NewRequest("GET", url, nil)
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		return nil, err
	}
	// Add Basic Auth header to the request
	auth.AddBasicAuth(req)

	// Declare buffer for the response body
	var bodyBuf bytes.Buffer
	// Declare response
	var res *http.Response
	// Determines whether or not the progress bar is diplayed or not
	if silent {
		// Send the request
		res, err = http.DefaultClient.Do(req)
		// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
		if err != nil {
			return nil, err
		}
		// Copy the HTTTP response body into the response body buffer
		io.Copy(io.Writer(&bodyBuf), res.Body)

	} else {
		// Create and Display the progress bar
		bar := progressbar.DefaultBytes(
			-1,
			"fetching",
		)
		// Send the request
		res, err = http.DefaultClient.Do(req)
		// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
		if err != nil {
			return nil, err
		}
		// Copy the HTTP response body into the progress bar and the response body buffer
		io.Copy(io.MultiWriter(bar, &bodyBuf), res.Body)
		// Inform the user that the download of the data is complete.
		println("done.")
	}
	// Defer the closing of the response body until the end of the function invocation.
	defer res.Body.Close()
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		return nil, err
	}
	// Parse the JSON string into an arbitrary golang map
	payload := map[string]interface{}{}
	err = json.Unmarshal(bodyBuf.Bytes(), &payload)
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		return nil, err
	}
	// If the HTTP status code is not 200(OK) then raise an error and return it with the reponse's error message as the go error message
	if res.StatusCode != 200 {
		return nil, errors.New(payload["error"].(string))
	}
	// Return the arbitrary map of the JSON response and a nil error pointer
	return &payload, nil
}

func GetTicket(ticketID int) (*map[string]interface{}, error) {
	url := urls.GetTicketURL(ticketID)
	// Create request struct
	req, err := http.NewRequest("GET", url, nil)
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		return nil, err
	}
	// Add Basic Auth header to the request
	auth.AddBasicAuth(req)

	// Declare buffer for the response body
	var bodyBuf bytes.Buffer
	// Declare response
	var res *http.Response
	// Send the request
	res, err = http.DefaultClient.Do(req)
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		return nil, err
	}
	// Create and Display the progress bar
	bar := progressbar.DefaultBytes(
		-1,
		"fetching",
	)
	// Send the request
	res, err = http.DefaultClient.Do(req)
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		return nil, err
	}
	// Copy the HTTP response body into the progress bar and the response body buffer
	io.Copy(io.MultiWriter(bar, &bodyBuf), res.Body)
	// Inform the user that the download of the data is complete.
	println("done.")
	// Defer the closing of the response body until the end of the function invocation.
	defer res.Body.Close()
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		return nil, err
	}
	// Parse the JSON string into an arbitrary golang map
	payload := map[string]interface{}{}
	err = json.Unmarshal(bodyBuf.Bytes(), &payload)
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		return nil, err
	}
	// If the HTTP status code is not 200(OK) then raise an error and return it with the reponse's error message as the go error message
	if res.StatusCode != 200 {
		return nil, errors.New(payload["error"].(string))
	}
	// Return the arbitrary map of the JSON response and a nil error pointer
	return &payload, nil
}

// Get the number of tickets the user has
func GetAccountTicketCount() (int, error) {
	// Create request struct
	req, err := http.NewRequest("GET", urls.COUNT_TICKETS, nil)
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		fmt.Println(&err)
	}
	// Add Basic Auth header to the request
	auth.AddBasicAuth(req)
	// Send the request
	res, err := http.DefaultClient.Do(req)
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		fmt.Println(&err)
	}
	// Defer the closing of the response body until the end of the function invocation.
	defer res.Body.Close()
	// Read the contents of the response body into a byte array
	body, err := io.ReadAll(res.Body)
	// Ensure there are no unchecked errors, if so, bubble them up to the CLI to be presented to the user
	if err != nil {
		return -1, err
	}
	// Parse the JSON string into an arbitrary golang map
	payload := map[string]interface{}{}
	json.Unmarshal(body, &payload)
	// If the HTTP status code is not 200(OK) then raise an error and return it with the reponse's error message as the go error message
	if res.StatusCode != 200 {
		return -1, errors.New(payload["error"].(string))
	}
	// Return the ticket count and a nil errror pointer
	return int(payload["count"].(map[string]interface{})["value"].(float64)), nil
}
