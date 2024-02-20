package model

import "lease/database"

type ProductImage struct {
	Model
	URL       string `gorm:"column:url;type:varchar(255);not null;comment:链接" json:"url"`
	ProductID int    `gorm:"column:product_id;type:int;not null;comment:商品ID" json:"product_id"`
}

func AddProductImage(productId int, path string) {
	database.DB.Create(&ProductImage{URL: path, ProductID: productId})
}

func GetProductImageList(productId int) (images []ProductImage) {
	database.DB.Where("product_id = ?", productId).Find(&images)
	return
}
