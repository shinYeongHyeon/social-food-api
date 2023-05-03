package shared_domain_test

import (
	"github.com/go-faker/faker/v4"
	"social-food-api/src/shared/domain"
	"testing"
)

func TestNameCreate(t *testing.T) {
	nameString := faker.Name()
	name, err := shared_domain.CreateName(nameString)

	if err != nil || name.Value() != nameString {
		t.Errorf("[Name] Create(%s) should return Name{%s}, but got error: %v", nameString, nameString, err)
	}
}

func TestNameCreateFailWithEmptyNameString(t *testing.T) {
	nameString := ""
	_, err := shared_domain.CreateName(nameString)

	if err == nil {
		t.Errorf("[Name] nameString should not be an empty string, but no error occured")
	}
}
