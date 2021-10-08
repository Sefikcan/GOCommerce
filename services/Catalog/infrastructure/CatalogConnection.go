package Connection

import (
	"catalog/common/constants"
	"catalog/infrastructure/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open(constants.PRODUCTCONNSTRING))
	if err != nil {
		panic("Could not connect to the database")
	}

	db.AutoMigrate(&entities.Product{})
	DB = db
}
