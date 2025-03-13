package models

type Certificate struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Course     string `json:"course"`
	IssuedTo   string `json:"issued_to"`
	IssueDate  string `json:"issue_date"`
	ExpiryDate string `json:"expiry_date"`
	Issuer     string `json:"issuer"`
	Content    string `json:"content"`
}
