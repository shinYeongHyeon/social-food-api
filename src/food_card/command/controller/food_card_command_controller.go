package food_card_command_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-food-api/src/food_card/command/application/create_food_card_use_case"
	"social-food-api/src/food_card/command/application/create_food_card_use_case/dto"
	"social-food-api/src/food_card/command/controller/dto"
	"social-food-api/src/food_card/repository"
	shared_core "social-food-api/src/shared/core"
	"social-food-api/src/shared/domain"
)

func CreateFoodCard(c *gin.Context) {
	var request food_card_command_controller_dto.CreateFoodCardRequest

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(shared_core.ThrowErrorWhenRequestBodyEmpty())
		return
	}

	if request.Name == "" || request.Address == "" || request.Latitude == 0 || request.Longitude == 0 {
		c.IndentedJSON(shared_core.ThrowErrorWhenSomeRequestValueNotExist())
		return
	}

	var name, nameError = shared_domain.CreateName(request.Name)
	var address, addressError = shared_domain.CreateAddress(request.Address)
	var latitude, latitudeError = shared_domain.CreateLatitudeLongitude(request.Latitude)
	var longitude, longitudeError = shared_domain.CreateLatitudeLongitude(request.Longitude)

	shared_core.ThrowErrorWhenCreatingValueObject(c, []error{nameError, addressError, latitudeError, longitudeError})

	var createFoodCardResponse = create_food_card_use_case.Execute(
		food_card_repository.GetRepository(),
		create_food_card_use_case_dto.CreateFoodCardUseCaseRequest{
			Name:      name,
			Address:   address,
			Latitude:  latitude,
			Longitude: longitude,
		},
	)

	if !createFoodCardResponse.IsSuccess || createFoodCardResponse.Code != "SUCCESS" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": "some error when create foodCard",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"code": "SUCCESS",
		"id":   createFoodCardResponse.FoodCard.GetUuid(),
	})
}
