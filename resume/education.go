package resume

type Education struct {
	Name         string `json:"name,omitempty"`
	Degree       string `json:"degree,omitempty"`
	FieldOfStudy string `json:"fieldOfStudy,omitempty"`
	StartDate    string `json:"startDate,omitempty"`
	EndDate      string `json:"endDate,omitempty"`
	URL          string `json:"url,omitempty"`
}
