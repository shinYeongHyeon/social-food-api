package food_card_repository

import (
	"gorm.io/gorm"
	"social-food-api/src/food_card/domain"
	"social-food-api/src/shared/infra"
	"social-food-api/src/shared/infra/entities"
)

// TODO: MaterializedView 와 Table 분리시, Repository Directory 도 command & query 로 변경해야 함

type FoodCardRepository struct {
	Repository *gorm.DB
}

var foodCardRepository *FoodCardRepository

func init() {
	if foodCardRepository != nil {
		return
	}

	foodCardRepository = &FoodCardRepository{shared_infra.GetManager().Db.Table(shared_infra_entities.FoodCardEntityTableName)}
}

// GetRepository : Get FoodCardRepository
func GetRepository() *FoodCardRepository {
	return foodCardRepository
}

// Create : create FoodCard Row
func (foodCardRepository *FoodCardRepository) Create(foodCard food_card_domain.FoodCard) error {
	// TODO: Research Create Error
	foodCardRepository.Repository.Create(&shared_infra_entities.FoodCardEntity{
		UUID:      foodCard.GetUuid(),
		Name:      foodCard.GetName().Value(),
		Address:   foodCard.GetAddress().Value(),
		Latitude:  foodCard.GetLatitude().Value(),
		Longitude: foodCard.GetLongitude().Value(),
		CreatedAt: foodCard.GetCreatedAt(),
		UpdatedAt: foodCard.GetUpdatedAt(),
	})

	return nil
}
