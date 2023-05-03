package shared_domain_test

import (
	"github.com/go-faker/faker/v4"
	"social-food-api/src/shared/domain"
	"testing"
)

func TestLatitudeLongitudeCreate(t *testing.T) {
	var testValue float64
	testValue = faker.Latitude()
	latitudeLongitude, err := shared_domain.CreateLatitudeLongitude(testValue)

	if err != nil || latitudeLongitude.Value() != testValue {
		t.Errorf("[LatitudeLongitude] Create(%f) should return LatitudeLongitude{%f}, but got error: %v", testValue, testValue, err)
	}
}

func TestLatitudeLongitudeCreateFailedWhenZero(t *testing.T) {
	_, err := shared_domain.CreateLatitudeLongitude(0)

	if err == nil {
		t.Errorf("[LatitudeLongitude] latitudeLongtidueValue should not be an zero, but no error occured")
	}
}
