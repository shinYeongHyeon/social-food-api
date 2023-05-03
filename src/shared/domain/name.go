package shared_domain

import "errors"

type Name struct {
	value string
}

func CreateName(nameString string) (Name, error) {
	if nameString == "" {
		return Name{}, errors.New("nameString must have a value")
	}

	return Name{nameString}, nil
}

func (name Name) Value() string {
	return name.value
}
