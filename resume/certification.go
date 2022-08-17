package resume

import "time"

type Certification struct {
	Name            string    `json:"name,omitempty"`
	Issuer          string    `json:"issuer,omitempty"`
	IssueDate       time.Time `json:"issueDate,omitempty"`
	ExpirationDate  time.Time `json:"expirationDate,omitempty"`
	VerificationURL string    `json:"verificationURL,omitempty"`
}
