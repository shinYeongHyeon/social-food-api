package shared_infra_entities

var FoodCardEntityTableName = "food_cards"

// FoodCardEntity : Entity of foodCard
type FoodCardEntity struct {
	UUID      string  `gorm:"primaryKey"`
	Name      string  `gorm:"type:varchar(255);not_null"`
	Address   string  `gorm:"type:varchar(255);not_null"`
	Latitude  float64 `gorm:"not_null"`
	Longitude float64 `gorm:"not_null"`
	CreatedAt int64   `gorm:"not_null"`
	UpdatedAt int64   `gorm:"not_null"`
}
