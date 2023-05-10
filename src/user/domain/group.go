package user_domain

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"social-food-api/src/shared/core"
	"social-food-api/src/shared/domain"
)

type Group struct {
	uuid      string
	ownerUuid string
	name      shared_domain.Name
	createdAt int64
	updatedAt int64
}

type NewGroupProps struct {
	OwnerUuid string
	Name      shared_domain.Name
}

func CreateNewGroup(props NewGroupProps) (Group, error) {
	nowSeconds := timestamppb.Now().Seconds

	return Group{
		uuid:      shared_core.CreateUuid(),
		ownerUuid: props.OwnerUuid,
		name:      props.Name,
		createdAt: nowSeconds,
		updatedAt: nowSeconds,
	}, nil
}

func (group Group) GetUuid() string {
	return group.uuid
}

func (group Group) GetOwnerUuid() string {
	return group.ownerUuid
}

func (group Group) GetName() shared_domain.Name {
	return group.name
}

func (group Group) GetCreatedAt() int64 {
	return group.createdAt
}

func (group Group) GetUpdatedAt() int64 {
	return group.updatedAt
}
