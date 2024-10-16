package models

import (
	"github.com/kamva/mgm/v3"
)

type Users struct {
	mgm.DefaultModel `bson:",inline"`
	Extension        string `json:"Extension" bson:"Extension"`
	Phone_Number     string `json:"Phone_Number" bson:"Phone_Number"`
	Country          string `json:"Country" bson:"Country"`
	Username         string `json:"Username" bson:"Username"`
	Description      string `json:"Description" bson:"Description"`
	// Profile_Picture  *multipart.FileHeader `json:"Profile_Picture" bson:"Profile_Picture"`
}

// func NewUser(extension string, phone_number string, country string, username string, profile_picture *multipart.FileHeader) *Users {
func NewUser(extension string, phone_number string, country string, username string, description string) *Users {
	return &Users{
		Extension:    extension,
		Phone_Number: phone_number,
		Country:      country,
		Username:     username,
		Description:  description,
		// Profile_Picture: profile_picture,
	}
}
