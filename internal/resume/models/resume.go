package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resume struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	FirstName      string             `json:"firstName" db:"firstName" bson:"firstName"`
	LastName       string             `json:"lastName" db:"lastName" bson:"lastName"`
	Title          string             `json:"title" db:"title" bson:"title"`
	Contacts       map[string]string  `json:"contacts" db:"contacts" bson:"contacts"`
	Summary        []string           `json:"summary" db:"summary" bson:"summary"`
	Skills         []SkillGroup       `json:"skills" db:"skills" bson:"skills"`
	Experience     []Experience       `json:"experience" db:"experience" bson:"experience"`
	Certifications []Certification    `json:"certifications" db:"certifications" bson:"certifications"`
	Education      []Education        `json:"education" db:"education" bson:"education"`
}

type SkillGroup struct {
	Name      string       `json:"name" db:"name" bson:"name"`
	SubGroups []SkillGroup `json:"subGroups,omitempty" db:"subGroups" bson:"subGroups"`
	Skills    []Skill      `json:"skills,omitempty" db:"skills" bson:"skills"`
}

type Skill struct {
	Name           string `json:"name" db:"name" bson:"name"`
	AdditionalInfo string `json:"additionalInfo,omitempty" db:"additionalInfo" bson:"additionalInfo"`
}

type Experience struct {
	Title     string    `json:"title" db:"title" bson:"title"`
	Company   string    `json:"company" db:"company" bson:"company"`
	StartDate time.Time `json:"startDate" db:"startDate" bson:"startDate"`
	EndDate   time.Time `json:"endDate" db:"endDate" bson:"endDate"`
	Projects  []Project `json:"projects,omitempty" db:"projects" bson:"projects"`
}

type Project struct {
	Client        string    `json:"client,omitempty" db:"client" bson:"client"`
	StartDate     time.Time `json:"startDate" db:"startDate" bson:"startDate"`
	EndDate       time.Time `json:"endDate" db:"endDate" bson:"endDate"`
	Roles         []string  `json:"projectRoles,omitempty" db:"projectRole" bson:"projectRole"`
	Envirnment    []string  `json:"envirnment,omitempty" db:"envirnment" bson:"envirnment"`
	TaskPerformed []string  `json:"taskPerformed,omitempty" db:"taskPerformed" bson:"taskPerformed"`
}

type Certification struct {
	Name            string    `json:"name" db:"name" bson:"name"`
	Issuer          string    `json:"issuer,omitempty" db:"issuer" bson:"issuer"`
	IssueDate       time.Time `json:"issueDate" db:"issueDate" bson:"issueDate"`
	ExpirationDate  time.Time `json:"expirationDate" db:"expirationDate" bson:"expirationDate"`
	VerificationURL string    `json:"verificationURL,omitempty" db:"verificationURL" bson:"verificationURL"`
}

type Education struct {
	Name         string    `json:"name" db:"name" bson:"name"`
	Degree       string    `json:"degree,omitempty" db:"degree" bson:"degree"`
	FieldOfStudy string    `json:"fieldOfStudy,omitempty" db:"fieldOfStudy" bson:"fieldOfStudy"`
	StartDate    time.Time `json:"startDate,omitempty" db:"startDate" bson:"startDate"`
	EndDate      time.Time `json:"endDate,omitempty" db:"endDate" bson:"endDate"`
	URL          string    `json:"url,omitempty" db:"url" bson:"url"`
}
