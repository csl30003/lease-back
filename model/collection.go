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

type CollectionDetail struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	MainImage string  `json:"main_image"`
	ProductID int     `json:"product_id"`
	UserID    int     `json:"user_id"`
}

func GetCollection(userId int) (collectionList []CollectionDetail) {
	database.DB.Table("collections").
		Select("collections.id, products.name, products.price, products.main_image, products.id as product_id, collections.user_id").
		Joins("left join products on collections.product_id = products.id AND products.deleted_at IS NULL").
		Where("collections.user_id = ? AND collections.deleted_at IS NULL", userId).
		Scan(&collectionList)
	return
}
