package shared_infra_entities

var UserEntityTableName = "users"

// UserEntity : Entity of user
type UserEntity struct {
	UUID      string `gorm:"primaryKey"`
	Id        string `gorm:"type:varchar(255);not_null;comment:유저 이메일"`
	Name      string `gorm:"type:varchar(255);not_null;comment:유저 닉네임"`
	UpdatedAt int64  `gorm:"not_null"`
	CreatedAt int64  `gorm:"not_null"`
}
