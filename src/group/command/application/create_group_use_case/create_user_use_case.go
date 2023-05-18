package create_group_use_case

import (
	"social-food-api/src/group/command/application/create_group_use_case/dto"
	"social-food-api/src/group/group_domain"
	"social-food-api/src/group/repository"
)

var failResponse = create_group_use_case_dto.CreateGroupUseCaseResponse{
	IsSuccess: false,
	Code:      "FAILED",
	Group:     group_domain.Group{},
}

func Execute(
	repository *group_repository.GroupRepository,
	request create_group_use_case_dto.CreateGroupUseCaseRequest,
) create_group_use_case_dto.CreateGroupUseCaseResponse {
	group, err := group_domain.CreateNewGroup(group_domain.NewGroupProps{
		OwnerUuid: request.OwnerUserId,
		Name:      request.Name,
	})

	if err != nil {
		return failResponse
	}

	saveErr := repository.Create(group)

	if saveErr != nil {
		return failResponse
	}

	return create_group_use_case_dto.CreateGroupUseCaseResponse{
		IsSuccess: true,
		Code:      "SUCCESS",
		Group:     group,
	}
}
