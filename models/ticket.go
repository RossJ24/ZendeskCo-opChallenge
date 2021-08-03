package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Go struct to represent the Ticket JSON
type Ticket struct {
	URL                 string        `json:"url"`
	ID                  int           `json:"id"`
	ExternalID          interface{}   `json:"external_id"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	Type                string        `json:"type"`
	Subject             string        `json:"subject"`
	RawSubject          string        `json:"raw_subject"`
	Description         string        `json:"description"`
	Priority            string        `json:"priority"`
	Status              string        `json:"status"`
	Recipient           interface{}   `json:"recipient"`
	RequesterID         int64         `json:"requester_id"`
	SubmitterID         int64         `json:"submitter_id"`
	AssigneeID          int64         `json:"assignee_id"`
	OrganizationID      interface{}   `json:"organization_id"`
	GroupID             int64         `json:"group_id"`
	CollaboratorIds     []interface{} `json:"collaborator_ids"`
	FollowerIds         []interface{} `json:"follower_ids"`
	EmailCcIds          []interface{} `json:"email_cc_ids"`
	ForumTopicID        interface{}   `json:"forum_topic_id"`
	ProblemID           interface{}   `json:"problem_id"`
	HasIncidents        bool          `json:"has_incidents"`
	IsPublic            bool          `json:"is_public"`
	DueAt               interface{}   `json:"due_at"`
	Tags                []string      `json:"tags"`
	CustomFields        []interface{} `json:"custom_fields"`
	SatisfactionRating  interface{}   `json:"satisfaction_rating"`
	SharingAgreementIds []interface{} `json:"sharing_agreement_ids"`
	Fields              []interface{} `json:"fields"`
	FollowupIds         []interface{} `json:"followup_ids"`
	TicketFormID        int64         `json:"ticket_form_id"`
	BrandID             int64         `json:"brand_id"`
	AllowChannelback    bool          `json:"allow_channelback"`
	AllowAttachments    bool          `json:"allow_attachments"`
}

// Conversts an arbitrary go map into a Ticket struct
func TicketFromJSON(jsonTicket interface{}) *Ticket {
	// Create an empty ticket
	ticket := &Ticket{}
	// Stringify the arbitary go map
	bytes, err := json.Marshal(jsonTicket)
	// If there is a parsing error return nil
	if err != nil {
		return nil
	}
	// Parse the json string into a Ticket struct
	err = json.Unmarshal(bytes, ticket)
	// If there is a parsing error return nil
	if err != nil {
		return nil
	}
	return ticket
}

// Get information from Ticket that is necessary for the List View.
func (ticket Ticket) GetListInfo() string {
	return fmt.Sprintf("ID: %4d  submitted by: %d status: '%s' subject: '%s'", ticket.ID, ticket.SubmitterID, ticket.Status, ticket.Subject)
}

// Get information from Ticket that is necessary for individual View.
func (ticket *Ticket) GetIndividualDisplayInfo() string {
	return fmt.Sprintf("Ticket %4d\n-------------\nsubmitted by: %d\nrequested by: %d\nstatus: '%s'\nsubject: '%s'\ncreated at: %s\nlast updated at: %s\nassigned to: %d\ndescription: '%s'", ticket.ID, ticket.SubmitterID, ticket.RequesterID, ticket.Status, ticket.Subject, ticket.CreatedAt.String(), ticket.UpdatedAt.String(), ticket.AssigneeID, ticket.Description)
}
