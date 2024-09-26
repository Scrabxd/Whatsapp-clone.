package models

import "github.com/kamva/mgm/v3"

type Users struct {
	mgm.DefaultModel `bson:",inline"`
	Extension        string `json:"Extension" bson:"Extension"`
	Phone_Number     string `json:"Phone_Number" bson:"Phone_Number"`
	Country          string `json:"Country" bson:"Country"`
}

func NewUser(extension string, phone_number string, country string) *Users {
	return &Users{
		Extension:    extension,
		Phone_Number: phone_number,
		Country:      country,
	}
}
