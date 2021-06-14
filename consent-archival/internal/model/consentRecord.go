package model

type consentRecord struct {
	UserId string `json:"userId"`
	Timestamp string `json:"acceptance_timestamp"`
	Attributes map[string]interface{} `json:"attributes"`
}