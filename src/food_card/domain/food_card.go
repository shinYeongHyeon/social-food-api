package food_card_domain

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"social-food-api/src/shared/core"
	"social-food-api/src/shared/domain"
)

type FoodCard struct {
	uuid      string
	name      shared_domain.Name
	address   shared_domain.Address
	latitude  shared_domain.LatitudeLongitude
	longitude shared_domain.LatitudeLongitude
	createdAt int64
	updatedAt int64
}

type NewFoodCardProps struct {
	Name      shared_domain.Name
	Address   shared_domain.Address
	Latitude  shared_domain.LatitudeLongitude
	Longitude shared_domain.LatitudeLongitude
}

func CreateNewFoodCard(props NewFoodCardProps) (FoodCard, error) {
	nowSeconds := timestamppb.Now().Seconds
	return FoodCard{
		uuid:      shared_core.CreateUuid(),
		name:      props.Name,
		address:   props.Address,
		latitude:  props.Latitude,
		longitude: props.Longitude,
		createdAt: nowSeconds,
		updatedAt: nowSeconds,
	}, nil
}

func (foodCard FoodCard) GetUuid() string {
	return foodCard.uuid
}

func (foodCard FoodCard) GetName() shared_domain.Name {
	return foodCard.name
}

func (foodCard FoodCard) GetAddress() shared_domain.Address {
	return foodCard.address
}

func (foodCard FoodCard) GetLatitude() shared_domain.LatitudeLongitude {
	return foodCard.latitude
}

func (foodCard FoodCard) GetLongitude() shared_domain.LatitudeLongitude {
	return foodCard.longitude
}

func (foodCard FoodCard) GetCreatedAt() int64 {
	return foodCard.createdAt
}

func (foodCard FoodCard) GetUpdatedAt() int64 {
	return foodCard.updatedAt
}
