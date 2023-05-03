package shared_domain

import "errors"

type Address struct {
	value string
}

func CreateAddress(addressString string) (Address, error) {
	if addressString == "" {
		return Address{}, errors.New("addressString must have value")
	}

	return Address{addressString}, nil
}

func (address Address) Value() string {
	return address.value
}
