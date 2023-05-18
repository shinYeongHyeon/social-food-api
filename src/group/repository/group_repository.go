package group_repository

import (
	"gorm.io/gorm"
	"social-food-api/src/group/group_domain"
	"social-food-api/src/shared/infra"
	"social-food-api/src/shared/infra/entities"
)

type GroupRepository struct {
	Repository *gorm.DB
}

var groupRepository *GroupRepository

func init() {
	if groupRepository != nil {
		return
	}

	groupRepository = &GroupRepository{shared_infra.GetManager().Db.Table(shared_infra_entities.GroupEntityTableName)}
}

func GetRepository() *GroupRepository {
	return groupRepository
}

func (r *GroupRepository) Create(group group_domain.Group) error {
	r.Repository.Create(&shared_infra_entities.GroupEntity{
		UUID:      group.GetUuid(),
		Name:      group.GetName().Value(),
		OwnerUuid: group.GetOwnerUuid(),
		CreatedAt: group.GetCreatedAt(),
		UpdatedAt: group.GetUpdatedAt(),
	})

	return nil
}
