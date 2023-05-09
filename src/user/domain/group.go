package user_domain

import (
	"social-food-api/src/shared/domain"
)

type Group struct {
	uuid      string
	ownerUuid string
	name      string
	createdAt int64
	updatedAt int64
}

type NewGroupProps struct {
	OwnerUuid string
	Name      shared_domain.Name
}
