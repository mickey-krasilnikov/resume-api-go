package resume

type Resume struct {
	FirstName      string            `json:"firstName"`
	LastName       string            `json:"lastName"`
	Title          string            `json:"title"`
	Contacts       map[string]string `json:"contacts"`
	Summary        []string          `json:"summary"`
	Skills         []SkillGroup      `json:"skills"`
	Experience     []Experience      `json:"experience"`
	Certifications []Certification   `json:"certifications"`
	Education      []Education       `json:"education"`
}
