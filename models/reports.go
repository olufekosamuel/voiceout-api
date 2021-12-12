package models

import (
	"github.com/Kamva/mgm/v2"
)

type Admin struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"first_name" bson:"first_name"`
	Surname          string `json:"surname" bson:"surname"`
	Email            string `json:"email" bson:"email"`
	PhoneNumber      string `json:"phone_number" bson:"phone_number"`
	Password         string `json:"password,omitempty" bson:"password"`
}

type Report struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Department       string `json:"department" bson:"department"`
	Description      string `json:"description" bson:"description"`
	File             string `json:"file" bson:"file"`
	Anonymous        bool   `json:"anonymous" bson:"anonymous"`
	Fullname         string `json:"full_name,omitempty" bson:"full_name"`
	Email            string `json:"email,omitempty" bson:"email"`
	Phone            string `json:"phone,omitempty" bson:"phone"`
}

type Feedback struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Description      string `json:"description" bson:"description"`
	Report           string `json:"report" bson:"report"`
}

type Department struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Department       string `json:"department" bson:"department"`
}

func NewAdmin() *Admin {
	return &Admin{}
}

func NewDepartment() *Department {
	return &Department{}
}

func NewReport() *Report {
	return &Report{}
}
