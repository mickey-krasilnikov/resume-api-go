package resume

import "time"

type Experience struct {
	Title         string    `json:"title,omitempty"`
	Company       string    `json:"company,omitempty"`
	Client        string    `json:"client,omitempty"`
	StartDate     time.Time `json:"startDate,omitempty"`
	EndDate       time.Time `json:"endDate,omitempty"`
	ProjectRole   string    `json:"projectRole,omitempty"`
	Envirnment    []string  `json:"envirnment,omitempty"`
	TaskPerformed []string  `json:"taskPerformed,omitempty"`
}
