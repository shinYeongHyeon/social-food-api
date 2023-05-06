package user_domain

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"social-food-api/src/shared/core"
	"social-food-api/src/shared/domain"
	"social-food-api/src/shared/enums"
)

type User struct {
	uuid      string
	userId    UserId
	name      shared_domain.Name
	from      shared_enums.LoginMethods
	createdAt int64
	updatedAt int64
}

type NewUserProps struct {
	UserId UserId
	Name   shared_domain.Name
	From   shared_enums.LoginMethods
}

func CreateNewUser(props NewUserProps) (User, error) {
	nowSeconds := timestamppb.Now().Seconds

	return User{
		uuid:      shared_core.CreateUuid(),
		userId:    props.UserId,
		name:      props.Name,
		from:      props.From,
		createdAt: nowSeconds,
		updatedAt: nowSeconds,
	}, nil
}

func (user User) GetUuid() string {
	return user.uuid
}

func (user User) GetUserId() UserId {
	return user.userId
}

func (user User) GetName() shared_domain.Name {
	return user.name
}

func (user User) GetFrom() shared_enums.LoginMethods {
	return user.from
}

func (user User) GetCreatedAt() int64 {
	return user.createdAt
}

func (user User) GetUpdatedAt() int64 {
	return user.updatedAt
}
