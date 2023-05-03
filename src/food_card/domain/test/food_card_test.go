package food_card_domain_test

import (
	"github.com/go-faker/faker/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
	"social-food-api/src/food_card/domain"
	"social-food-api/src/shared/domain"
	"testing"
)

func TestNewFoodCardCreate(t *testing.T) {
	nameString := faker.Name()
	addressString := faker.Sentence()
	latitudeValue := faker.Latitude()
	longitudeValue := faker.Longitude()

	name, _ := shared_domain.CreateName(nameString)
	address, _ := shared_domain.CreateAddress(addressString)
	latitude, _ := shared_domain.CreateLatitudeLongitude(latitudeValue)
	longitude, _ := shared_domain.CreateLatitudeLongitude(longitudeValue)

	foodCard, err := food_card_domain.CreateNewFoodCard(food_card_domain.NewFoodCardProps{
		Name:      name,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	})

	if err != nil ||
		foodCard.GetName().Value() != nameString ||
		foodCard.GetAddress().Value() != addressString ||
		foodCard.GetLatitude().Value() != latitudeValue ||
		foodCard.GetLongitude().Value() != longitudeValue ||
		foodCard.GetCreatedAt() > timestamppb.Now().Seconds ||
		foodCard.GetUpdatedAt() > timestamppb.Now().Seconds ||
		foodCard.GetCreatedAt() != foodCard.GetUpdatedAt() {
		t.Errorf("[FoodCard] FoodCard Created Failed : %v, got error: %v", foodCard, err)
	}
}
