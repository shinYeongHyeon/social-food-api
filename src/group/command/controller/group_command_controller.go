package group_command_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-food-api/src/group/command/application/create_group_use_case"
	"social-food-api/src/group/command/application/create_group_use_case/dto"
	"social-food-api/src/group/command/controller/dto"
	"social-food-api/src/group/repository"
	"social-food-api/src/shared/core"
	"social-food-api/src/shared/domain"
)

func CreateGroup(c *gin.Context) {
	isAuthorized, _ := c.Get(shared_core.IsAuthorized)
	userId, isUserIdExist := c.Get(shared_core.Authorized)

	if !(isAuthorized.(bool)) || !isUserIdExist || userId == nil || userId == "" {
		c.IndentedJSON(shared_core.ThrowUnauthorized())
		return
	}

	var request group_command_controller_dto.CreateGroupRequest

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(shared_core.ThrowErrorWhenRequestBodyEmpty())
		return
	}

	groupName, groupNameError := shared_domain.CreateName(request.Name)

	hasError, voErr := shared_core.ValidateValueObjectErrors([]error{groupNameError})

	if hasError {
		c.IndentedJSON(shared_core.ThrowErrorWhenCreatingValueObject(voErr))
		return
	}

	var createGroupResponse = create_group_use_case.Execute(
		group_repository.GetRepository(),
		create_group_use_case_dto.CreateGroupUseCaseRequest{
			OwnerUserId: userId.(string),
			Name:        groupName,
		},
	)

	if !createGroupResponse.IsSuccess || createGroupResponse.Code != "SUCCESS" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": "some error when create group",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"code": "SUCCESS",
		"id":   createGroupResponse.Group.GetUuid(),
	})
}
