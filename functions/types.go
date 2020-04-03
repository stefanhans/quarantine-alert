package functions

import "time"

type Registration struct {
	ReporterID string    `json:"reporter"`
	Contagious bool      `json:"contagious"`
	Since      time.Time `json:"time-contagion-updated"`
}

type Contact struct {
	ReporterID string `json:"reporter"`
	Opposite
}

type Opposite struct {
	ContactID   string    `json:"contact"`
	ContactTime time.Time `json:"contact-time"`
}

// Represents a flat Firestore collection
type ContactsCollection map[string]map[string]interface{}

type Report struct {
	RequesterID string   `json:"requester"`
	Report      []string `json:"report"`
}

type Format struct {
	Indent bool `json:"indent"`
}

var (
	projectID = "quarantine-alert-22365"

	daysOfBeforeContagionPeriod = 12
	daysOfAfterContagionPeriod  = 12
)
