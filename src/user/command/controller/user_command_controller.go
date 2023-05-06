package user_command_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-food-api/src/shared/core"
	"social-food-api/src/shared/domain"
	shared_enums "social-food-api/src/shared/enums"
	"social-food-api/src/user/command/application/create_user_use_case"
	create_user_use_case_dto "social-food-api/src/user/command/application/create_user_use_case/dto"
	"social-food-api/src/user/command/controller/dto"
	"social-food-api/src/user/domain"
	user_repository "social-food-api/src/user/repository"
)

func CreateUser(c *gin.Context) {
	var request user_command_controller_dto.CreateUserRequest

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(shared_core.ThrowErrorWhenRequestBodyEmpty())
		return
	}

	if request.Id == "" || request.Name == "" || request.LoginMethod == "" {
		c.IndentedJSON(shared_core.ThrowErrorWhenSomeRequestValueNotExist())
		return
	}

	userName, userNameError := shared_domain.CreateName(request.Name)
	userId, userIdError := user_domain.CreateUserId(request.Id)
	loginMethod, loginMethodError := shared_enums.FromString(request.LoginMethod)

	shared_core.ThrowErrorWhenCreatingValueObject(c, []error{userNameError, userIdError, loginMethodError})

	var createUserResponse = create_user_use_case.Execute(
		user_repository.GetRepository(),
		create_user_use_case_dto.CreateUserUseCaseRequest{
			UserId:      userId,
			Name:        userName,
			LoginMethod: loginMethod,
		},
	)

	if !createUserResponse.IsSuccess || createUserResponse.Code != "SUCCESS" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": "some error when create user",
		})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"code": "SUCCESS",
		"id":   createUserResponse.User.GetUuid(),
	})
}
