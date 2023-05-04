package create_food_card_use_case

import (
	"social-food-api/src/food_card/command/application/create_food_card_use_case/dto"
	"social-food-api/src/food_card/domain"
	food_card_repository "social-food-api/src/food_card/repository"
)

func Execute(
	repository *food_card_repository.FoodCardRepository,
	request create_food_card_use_case_dto.CreateFoodCardUseCaseRequest,
) create_food_card_use_case_dto.CreateFoodCardUseCaseResponse {
	foodCard, err := food_card_domain.CreateNewFoodCard(food_card_domain.NewFoodCardProps{
		Name:      request.Name,
		Address:   request.Address,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
	})

	if err != nil {
		return create_food_card_use_case_dto.CreateFoodCardUseCaseResponse{
			IsSuccess: false,
			Code:      "FAILED",
			FoodCard:  food_card_domain.FoodCard{},
		}
	}

	saveErr := repository.Create(foodCard)

	if saveErr != nil {
		return create_food_card_use_case_dto.CreateFoodCardUseCaseResponse{
			IsSuccess: false,
			Code:      "FAILED",
			FoodCard:  food_card_domain.FoodCard{},
		}
	}

	return create_food_card_use_case_dto.CreateFoodCardUseCaseResponse{
		IsSuccess: true,
		Code:      "SUCCESS",
		FoodCard:  foodCard,
	}
}
