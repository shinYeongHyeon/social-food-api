package shared_domain_test

import (
	"github.com/go-faker/faker/v4"
	"social-food-api/src/shared/domain"
	"testing"
)

func TestAddressCreate(t *testing.T) {
	addressString := faker.Sentence()
	address, err := shared_domain.CreateAddress(addressString)

	if err != nil || address.Value() != addressString {
		t.Errorf("[Address] CreateAddress(%s) should return Address{%s}, but got error: %v", addressString, addressString, err)
	}
}

func TestAddressCreateFailWithEmptyNameString(t *testing.T) {
	addressString := ""
	_, err := shared_domain.CreateAddress(addressString)

	if err == nil {
		t.Errorf("[Address] addressString should not be an empty string, but no error occured")
	}
}
