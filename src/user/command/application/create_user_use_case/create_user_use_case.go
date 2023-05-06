package create_user_use_case

import (
	"social-food-api/src/user/command/application/create_user_use_case/dto"
	"social-food-api/src/user/domain"
	"social-food-api/src/user/repository"
)

func Execute(
	repository *user_repository.UserRepository,
	request create_user_use_case_dto.CreateUserUseCaseRequest,
) create_user_use_case_dto.CreateUserUseCaseResponse {
	user, err := user_domain.CreateNewUser(user_domain.NewUserProps{
		UserId: request.UserId,
		Name:   request.Name,
		From:   request.LoginMethod,
	})

	if err != nil {
		return create_user_use_case_dto.CreateUserUseCaseResponse{
			IsSuccess: false,
			Code:      "FAILED",
			User:      user_domain.User{},
		}
	}

	saveErr := repository.Create(user)

	if saveErr != nil {
		return create_user_use_case_dto.CreateUserUseCaseResponse{
			IsSuccess: false,
			Code:      "FAILED",
			User:      user_domain.User{},
		}
	}

	return create_user_use_case_dto.CreateUserUseCaseResponse{
		IsSuccess: true,
		Code:      "SUCCESS",
		User:      user,
	}
}
