package parking

import (
	"strings"
)

type Vehicle struct {
	registration string
}

func NewVehicle(registration string) (Vehicle, error) {
	registration = strings.ToUpper(strings.TrimSpace(registration))

	if registration == "" {
		return Vehicle{}, ErrInvalidVehicle
	}

	return Vehicle{registration: registration}, nil
}

func (v Vehicle) Registration() string {
	return v.registration
}
