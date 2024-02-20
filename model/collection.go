package model

import (
	"lease/database"
)

type Collection struct {
	Model
	ProductID int `gorm:"column:product_id;type:int;not null;comment:商品ID" json:"product_id"`
	UserID    int `gorm:"column:user_id;type:int;not null;comment:用户ID" json:"user_id"`
}

func ExistCollection(productID, userID int) int {
	var collection Collection
	database.DB.Where("product_id = ? AND user_id = ?", productID, userID).First(&collection)

	return collection.ID
}

func DeleteCollection(collection Collection) {
	database.DB.Delete(&collection)
}

func AddCollection(collection Collection) {
	database.DB.Create(&collection)
}
