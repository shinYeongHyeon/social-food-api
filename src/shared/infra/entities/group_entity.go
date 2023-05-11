package shared_infra_entities

var GroupEntityTableName = "groups"

// GroupEntity : Entity of group
type GroupEntity struct {
	UUID      string `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255);not_null"`
	OwnerUuid string `gorm:"type:varchar(255);not_null"`
	CreatedAt int64  `gorm:"not_null"`
	UpdatedAt int64  `gorm:"not_null"`
}
