package Helpers

import (
	types "WhatsAppClone/Types"
	"encoding/json"
)

func Parser(body []byte) (*types.UserPhone, error) {

	var phone_numbers types.UserPhone
	if err := json.Unmarshal(body, &phone_numbers); err != nil {
		return nil, err
	}

	return &phone_numbers, nil
}
