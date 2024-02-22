package dto

import (
	"gorm.io/gorm"
	"lease/model"
	"time"
)

type GetProductResp struct {
	ID              int                  `json:"id"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
	DeletedAt       gorm.DeletedAt       `json:"deleted_at"`
	Name            string               `json:"name"`
	Price           float64              `json:"price"`
	Detail          string               `json:"detail"`
	MainImage       string               `json:"main_image"`
	Stock           int                  `json:"stock"`
	Delivery        int                  `json:"delivery"`
	Freight         float64              `json:"freight"`
	Fineness        int                  `json:"fineness"`
	UsedYears       int                  `json:"used_years"`
	Status          int                  `json:"status"`
	CategoryID      int                  `json:"category_id"`
	CategoryName    string               `json:"category_name"`
	AddressID       int                  `json:"address_id"`
	AddressName     string               `json:"address_name"`
	AddressPhone    string               `json:"address_phone"`
	AddressProvince string               `json:"address_province"`
	AddressCity     string               `json:"address_city"`
	AddressDistrict string               `json:"address_district"`
	AddressDetail   string               `json:"address_detail"`
	UserID          int                  `json:"user_id"`
	UserName        string               `json:"user_name"`
	UserAvatar      string               `json:"user_avatar"`
	ImageList       []model.ProductImage `json:"image_list"`
}
