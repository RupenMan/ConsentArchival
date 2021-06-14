package model

type consentRecord struct {
	ProfileReferenceId string `json:"profileReferenceId"`
	Timestamp string `json:"acceptance_timestamp"`
	Attributes map[string]interface{} `json:"attributes"`
}