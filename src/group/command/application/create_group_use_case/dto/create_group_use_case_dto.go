package create_group_use_case_dto

import (
	"social-food-api/src/group/group_domain"
	"social-food-api/src/shared/domain"
)

type CreateGroupUseCaseRequest struct {
	OwnerUserId string
	Name        shared_domain.Name
}

type CreateGroupUseCaseResponse struct {
	IsSuccess bool
	Code      string
	Group     group_domain.Group
}
