package resume

type SkillGroup struct {
	Name      string       `json:"name,omitempty"`
	SubGroups []SkillGroup `json:"subGroups,omitempty"`
	Skills    []Skill      `json:"skills,omitempty"`
}

type Skill struct {
	Name           string `json:"name,omitempty"`
	AdditionalInfo string `json:"additionalInfo,omitempty"`
}
