package user_repository

import (
	"gorm.io/gorm"
	"social-food-api/src/shared/infra"
	"social-food-api/src/shared/infra/entities"
	"social-food-api/src/user/domain"
)

type UserRepository struct {
	Repository *gorm.DB
}

var userRepository *UserRepository

func init() {
	if userRepository != nil {
		return
	}

	userRepository = &UserRepository{shared_infra.GetManager().Db.Table(shared_infra_entities.UserEntityTableName)}
}

func GetRepository() *UserRepository {
	return userRepository
}

func (r *UserRepository) Create(user user_domain.User) error {
	r.Repository.Create(&shared_infra_entities.UserEntity{
		UUID:      user.GetUuid(),
		Id:        user.GetUserId().Value(),
		Name:      user.GetName().Value(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	})

	return nil
}
