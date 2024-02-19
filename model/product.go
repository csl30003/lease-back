package model

import "lease/database"

type Product struct {
	Model
	Name       string  `gorm:"column:name;type:varchar(255);not null;comment:商品名称" json:"name"`
	Price      float64 `gorm:"column:price;type:decimal(10,2);not null;comment:商品价格" json:"price"`
	Detail     string  `gorm:"column:detail;type:text;default:null;comment:商品详情" json:"detail"`
	MainImage  string  `gorm:"column:main_image;type:varchar(255);default:null;comment:商品主图" json:"main_image"`
	Stock      int     `gorm:"column:stock;type:int;not null;comment:商品库存" json:"stock"`
	Delivery   int     `gorm:"column:delivery;type:int;default:0;comment:配送方式，0邮寄、1自提" json:"delivery"`
	Freight    float64 `gorm:"column:freight;type:decimal(10,2);default:0;comment:运费" json:"freight"`
	Condition  int     `gorm:"column:condition;type:int;default:0;comment:成色，0全新、1几乎全新、2轻微使用痕迹、3明显使用痕迹" json:"condition"`
	UsedYears  int     `gorm:"column:used_years;type:int;default:0;comment:已用年限" json:"used_years"`
	Status     int     `gorm:"column:status;type:tinyint(1);default:0;comment:状态，0未发布，1已发布，2已下架" json:"status"`
	CategoryID int     `gorm:"column:category_id;type:int;not null;comment:商品分类ID" json:"category_id"`
	AddressID  int     `gorm:"column:address_id;type:int;not null;comment:地址ID" json:"address_id"`
	UserID     int     `gorm:"column:user_id;type:int;not null;comment:用户ID" json:"user_id"`
}

func AddProduct(product Product) (id int) {
	database.DB.Create(&product)
	id = product.ID
	return
}

func UpdateProductMainImage(id int, mainImage string) {
	database.DB.Model(&Product{}).Where("id = ?", id).Update("main_image", mainImage)
}

func UpdateProductStatus(product Product) {
	database.DB.Model(&Product{}).Where("id = ?", product.ID).Update("status", product.Status)
}

func GetMyProduct(userId, status int) (products []Product) {
	database.DB.Where("user_id = ? and status = ?", userId, status).Find(&products)
	return
}
