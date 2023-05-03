package shared_domain

import "errors"

type LatitudeLongitude struct {
	value float64
}

func CreateLatitudeLongitude(value float64) (LatitudeLongitude, error) {
	if value == 0 {
		return LatitudeLongitude{}, errors.New("latitudeLongitude value must be not zero")
	}

	return LatitudeLongitude{value}, nil
}

func (latitudeLongitude LatitudeLongitude) Value() float64 {
	return latitudeLongitude.value
}
