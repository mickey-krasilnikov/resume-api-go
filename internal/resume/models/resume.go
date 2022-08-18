package models

import "time"

type Resume struct {
	ID             int               `json:"id" db:"id" bson:"_id"`
	FirstName      string            `json:"firstName" db:"firstName" bson:"firstName"`
	LastName       string            `json:"lastName" db:"lastName" bson:"lastName"`
	Title          string            `json:"title" db:"title" bson:"title"`
	Contacts       map[string]string `json:"contacts" db:"contacts" bson:"contacts"`
	Summary        []string          `json:"summary" db:"summary" bson:"summary"`
	Skills         []SkillGroup      `json:"skills" db:"skills" bson:"skills"`
	Experience     []Experience      `json:"experience" db:"experience" bson:"experience"`
	Certifications []Certification   `json:"certifications" db:"certifications" bson:"certifications"`
	Education      []Education       `json:"education" db:"education" bson:"education"`
}

type SkillGroup struct {
	ID        int          `json:"id" db:"id" bson:"_id"`
	Name      string       `json:"name,omitempty" db:"name" bson:"name"`
	SubGroups []SkillGroup `json:"subGroups,omitempty" db:"subGroups" bson:"subGroups"`
	Skills    []Skill      `json:"skills,omitempty" db:"skills" bson:"skills"`
}

type Skill struct {
	ID             int    `json:"id" db:"id" bson:"_id"`
	Name           string `json:"name,omitempty" db:"name" bson:"name"`
	AdditionalInfo string `json:"additionalInfo,omitempty" db:"additionalInfo" bson:"additionalInfo"`
}

type Experience struct {
	ID            int       `json:"id" db:"id" bson:"_id"`
	Title         string    `json:"title,omitempty" db:"title" bson:"title"`
	Company       string    `json:"company,omitempty" db:"company" bson:"company"`
	Client        string    `json:"client,omitempty" db:"client" bson:"client"`
	StartDate     time.Time `json:"startDate,omitempty" db:"startDate" bson:"startDate"`
	EndDate       time.Time `json:"endDate,omitempty" db:"endDate" bson:"endDate"`
	ProjectRole   string    `json:"projectRole,omitempty" db:"projectRole" bson:"projectRole"`
	Envirnment    []string  `json:"envirnment,omitempty" db:"envirnment" bson:"envirnment"`
	TaskPerformed []string  `json:"taskPerformed,omitempty" db:"taskPerformed" bson:"taskPerformed"`
}

type Certification struct {
	ID              int       `json:"id" db:"id" bson:"_id"`
	Name            string    `json:"name,omitempty" db:"name" bson:"name"`
	Issuer          string    `json:"issuer,omitempty" db:"issuer" bson:"issuer"`
	IssueDate       time.Time `json:"issueDate,omitempty" db:"issueDate" bson:"issueDate"`
	ExpirationDate  time.Time `json:"expirationDate,omitempty" db:"expirationDate" bson:"expirationDate"`
	VerificationURL string    `json:"verificationURL,omitempty" db:"verificationURL" bson:"verificationURL"`
}

type Education struct {
	ID           int       `json:"id" db:"id" bson:"_id"`
	Name         string    `json:"name,omitempty" db:"name" bson:"name"`
	Degree       string    `json:"degree,omitempty" db:"degree" bson:"degree"`
	FieldOfStudy string    `json:"fieldOfStudy,omitempty" db:"fieldOfStudy" bson:"fieldOfStudy"`
	StartDate    time.Time `json:"startDate,omitempty" db:"startDate" bson:"startDate"`
	EndDate      time.Time `json:"endDate,omitempty" db:"endDate" bson:"endDate"`
	URL          string    `json:"url,omitempty" db:"url" bson:"url"`
}
