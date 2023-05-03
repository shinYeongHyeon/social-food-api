package create_food_card_use_case_dto

import (
	"social-food-api/src/food_card/domain"
	"social-food-api/src/shared/domain"
)

type CreateFoodCardUseCaseRequest struct {
	Name      shared_domain.Name
	Address   shared_domain.Address
	Latitude  shared_domain.LatitudeLongitude
	Longitude shared_domain.LatitudeLongitude
}

type CreateFoodCardUseCaseResponse struct {
	IsSuccess bool
	Code      string
	FoodCard  food_card_domain.FoodCard
}
