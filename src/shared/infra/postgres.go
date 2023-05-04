package shared_infra

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	shared_infra_entities "social-food-api/src/shared/infra/entities"
)

type PostgresManager struct {
	Db *gorm.DB
}

var Manager *PostgresManager

func init() {
	log.Println("PostgresManager init")

	if Manager != nil {
		return
	}

	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connect Error !")
	}

	Manager = &PostgresManager{Db: db}

	fatalErrorWhenAutoMigrating(Manager.Db.Table(shared_infra_entities.FoodCardEntityTableName).AutoMigrate(&shared_infra_entities.FoodCardEntity{}))
}

// GetManager : Get PostgresManager
func GetManager() *PostgresManager {
	return Manager
}

func fatalErrorWhenAutoMigrating(err error) {
	if err != nil {
		log.Fatalln("Error occurred when autoMigration : ", err)
	}

	return
}
