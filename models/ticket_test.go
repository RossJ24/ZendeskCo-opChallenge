package models

import (
	"reflect"
	"testing"
	"time"
)

// Tests that tickets are converted to JSON properly
func TestTicketToJSON(t *testing.T) {
	// Ticket representation in an arbitaary go map
	jsonTestRep := map[string]interface{}{
		"url":         "https://zccrossj.zendesk.com/api/v2/tickets/1.json",
		"id":          1,
		"external_id": nil,
		"via": map[string]interface{}{
			"channel": "sample_ticket",
			"source": map[string]interface{}{
				"from": map[string]interface{}{},
				"to":   map[string]interface{}{},
				"rel":  nil,
			},
		},
		"created_at":       "2021-07-28T19:35:09Z",
		"updated_at":       "2021-07-28T19:35:10Z",
		"type":             "incident",
		"subject":          "Sample ticket: Meet the ticket",
		"raw_subject":      "Sample ticket: Meet the ticket",
		"description":      "Hi Ross,\n\nThis is your first ticket. Ta-da! Any customer request sent to your supported channels (email, chat, voicemail, web form, and tweet) will become a Support ticket, just like this one. Respond to this ticket by typing a message and clicking Submit. You can also see how an email becomes a ticket by emailing your new account, support@zccrossj.zendesk.com. Your ticket will appear in ticket views.\n\nThat's the ticket on tickets. If you want to learn more, check out: \nhttps://support.zendesk.com/hc/en-us/articles/203691476\n",
		"priority":         "normal",
		"status":           "open",
		"recipient":        nil,
		"requester_id":     418757224432,
		"submitter_id":     418757209152,
		"assignee_id":      418757209152,
		"organization_id":  nil,
		"group_id":         360020699572,
		"collaborator_ids": []interface{}{},
		"follower_ids":     []interface{}{},
		"email_cc_ids":     []interface{}{},
		"forum_topic_id":   nil,
		"problem_id":       nil,
		"has_incidents":    false,
		"is_public":        true,
		"due_at":           nil,
		"tags": []interface{}{
			"sample",
			"support",
			"zendesk",
		},
		"custom_fields":         []interface{}{},
		"satisfaction_rating":   nil,
		"sharing_agreement_ids": []interface{}{},
		"fields":                []interface{}{},
		"followup_ids":          []interface{}{},
		"ticket_form_id":        360003057492,
		"brand_id":              360006791832,
		"allow_channelback":     false,
		"allow_attachments":     true,
	}
	outputTicket := TicketFromJSON(jsonTestRep)

	timestamp1, err := time.Parse(time.RFC3339, "2021-07-28T19:35:09Z")
	if err != nil {
		t.Errorf("Unable to parse timestamp: %s", err.Error())
	}
	timestamp2, err := time.Parse(time.RFC3339, "2021-07-28T19:35:10Z")
	if err != nil {
		t.Errorf("Unable to parse timestamp: %s", err.Error())
	}
	// Expected ticket value
	expectedTicket := Ticket{
		"https://zccrossj.zendesk.com/api/v2/tickets/1.json",
		1,
		nil,
		timestamp1,
		timestamp2,
		"incident",
		"Sample ticket: Meet the ticket",
		"Sample ticket: Meet the ticket",
		"Hi Ross,\n\nThis is your first ticket. Ta-da! Any customer request sent to your supported channels (email, chat, voicemail, web form, and tweet) will become a Support ticket, just like this one. Respond to this ticket by typing a message and clicking Submit. You can also see how an email becomes a ticket by emailing your new account, support@zccrossj.zendesk.com. Your ticket will appear in ticket views.\n\nThat's the ticket on tickets. If you want to learn more, check out: \nhttps://support.zendesk.com/hc/en-us/articles/203691476\n",
		"normal",
		"open",
		nil,
		418757224432,
		418757209152,
		418757209152,
		nil,
		360020699572,
		[]interface{}{},
		[]interface{}{},
		[]interface{}{},
		nil,
		nil,
		false,
		true,
		nil,
		[]string{
			"sample",
			"support",
			"zendesk",
		},
		[]interface{}{},
		nil,
		[]interface{}{},
		[]interface{}{},
		[]interface{}{},
		360003057492,
		360006791832,
		false,
		true,
	}
	// Use deep reflection to ensure equality between the fields of each of the structs
	if !reflect.DeepEqual(*outputTicket, expectedTicket) {
		t.Log("output: \n")
		t.Log(*outputTicket)
		t.Log("expected:\n")
		t.Log(expectedTicket)
		t.Errorf("Output did not equal expected.")
	}
}

func TestGetListInfo(t *testing.T) {
	// Expected string
	expected := "ID:    1  submitted by: 418757209152 status: 'open' subject: 'Sample ticket: Meet the ticket'"

	// Create ticket for testing
	timestamp1, err := time.Parse(time.RFC3339, "2021-07-28T19:35:09Z")
	if err != nil {
		t.Errorf("Unable to parse timestamp: %s", err.Error())
	}
	timestamp2, err := time.Parse(time.RFC3339, "2021-07-28T19:35:10Z")
	if err != nil {
		t.Errorf("Unable to parse timestamp: %s", err.Error())
	}
	testTicket := Ticket{
		"https://zccrossj.zendesk.com/api/v2/tickets/1.json",
		1,
		nil,
		timestamp1,
		timestamp2,
		"incident",
		"Sample ticket: Meet the ticket",
		"Sample ticket: Meet the ticket",
		"Hi Ross,\n\nThis is your first ticket. Ta-da! Any customer request sent to your supported channels (email, chat, voicemail, web form, and tweet) will become a Support ticket, just like this one. Respond to this ticket by typing a message and clicking Submit. You can also see how an email becomes a ticket by emailing your new account, support@zccrossj.zendesk.com. Your ticket will appear in ticket views.\n\nThat's the ticket on tickets. If you want to learn more, check out: \nhttps://support.zendesk.com/hc/en-us/articles/203691476\n",
		"normal",
		"open",
		nil,
		418757224432,
		418757209152,
		418757209152,
		nil,
		360020699572,
		[]interface{}{},
		[]interface{}{},
		[]interface{}{},
		nil,
		nil,
		false,
		true,
		nil,
		[]string{
			"sample",
			"support",
			"zendesk",
		},
		[]interface{}{},
		nil,
		[]interface{}{},
		[]interface{}{},
		[]interface{}{},
		360003057492,
		360006791832,
		false,
		true,
	}
	// Get the output of the list display information of the test ticket
	output := testTicket.GetListInfo()
	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}
}

func TestGetIndividualDisplayInfo(t *testing.T) {
	// Expected string
	expected := "Ticket    1\n-------------\nsubmitted by: 418757209152\nrequested by: 418757224432\nstatus: 'open'\nsubject: 'Sample ticket: Meet the ticket'\ncreated at: 2021-07-28 19:35:09 +0000 UTC\nlast updated at: 2021-07-28 19:35:10 +0000 UTC\nassigned to: 418757209152\ndescription: 'Hi Ross,\n\nThis is your first ticket. Ta-da! Any customer request sent to your supported channels (email, chat, voicemail, web form, and tweet) will become a Support ticket, just like this one. Respond to this ticket by typing a message and clicking Submit. You can also see how an email becomes a ticket by emailing your new account, support@zccrossj.zendesk.com. Your ticket will appear in ticket views.\n\nThat's the ticket on tickets. If you want to learn more, check out: \nhttps://support.zendesk.com/hc/en-us/articles/203691476\n'"

	// Create ticket for testing
	timestamp1, err := time.Parse(time.RFC3339, "2021-07-28T19:35:09Z")
	if err != nil {
		t.Errorf("Unable to parse timestamp: %s", err.Error())
	}
	timestamp2, err := time.Parse(time.RFC3339, "2021-07-28T19:35:10Z")
	if err != nil {
		t.Errorf("Unable to parse timestamp: %s", err.Error())
	}
	testTicket := Ticket{
		"https://zccrossj.zendesk.com/api/v2/tickets/1.json",
		1,
		nil,
		timestamp1,
		timestamp2,
		"incident",
		"Sample ticket: Meet the ticket",
		"Sample ticket: Meet the ticket",
		"Hi Ross,\n\nThis is your first ticket. Ta-da! Any customer request sent to your supported channels (email, chat, voicemail, web form, and tweet) will become a Support ticket, just like this one. Respond to this ticket by typing a message and clicking Submit. You can also see how an email becomes a ticket by emailing your new account, support@zccrossj.zendesk.com. Your ticket will appear in ticket views.\n\nThat's the ticket on tickets. If you want to learn more, check out: \nhttps://support.zendesk.com/hc/en-us/articles/203691476\n",
		"normal",
		"open",
		nil,
		418757224432,
		418757209152,
		418757209152,
		nil,
		360020699572,
		[]interface{}{},
		[]interface{}{},
		[]interface{}{},
		nil,
		nil,
		false,
		true,
		nil,
		[]string{
			"sample",
			"support",
			"zendesk",
		},
		[]interface{}{},
		nil,
		[]interface{}{},
		[]interface{}{},
		[]interface{}{},
		360003057492,
		360006791832,
		false,
		true,
	}
	// Get the output of the individual display information of the test ticket
	output := testTicket.GetIndividualDisplayInfo()
	if output != expected {
		t.Errorf("Output: %s, did not equal expected: %s.", output, expected)
	}

}
