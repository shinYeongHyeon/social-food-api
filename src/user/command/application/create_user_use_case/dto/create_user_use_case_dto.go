package create_user_use_case_dto

import (
	"social-food-api/src/shared/domain"
	shared_enums "social-food-api/src/shared/enums"
	"social-food-api/src/user/domain"
)

type CreateUserUseCaseRequest struct {
	UserId      user_domain.UserId
	Name        shared_domain.Name
	LoginMethod shared_enums.LoginMethods
}

type CreateUserUseCaseResponse struct {
	IsSuccess bool
	Code      string
	User      user_domain.User
}
