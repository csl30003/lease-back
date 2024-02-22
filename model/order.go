package model

import (
	"lease/database"
	"time"
)

type Order struct {
	Model
	Identifier          string    `gorm:"type:varchar(24);default:null;comment:'订单编号'" json:"identifier"`
	Status              int       `gorm:"type:tinyint(1);not null;default:1;comment:'状态 0已取消 1未付款 2已付款 3已发货 4已收货 5归还中 6已归还 7待解决，8已完成'" json:"status"`
	PayTime             time.Time `gorm:"type:datetime;default:null;comment:'我方支付时间'" json:"pay_time"`
	HisDeliveryTime     time.Time `gorm:"type:datetime;default:null;comment:'对方发货时间'" json:"his_delivery_time"`
	MyReceiveTime       time.Time `gorm:"type:datetime;default:null;comment:'我方收货时间'" json:"my_receive_time"`
	ReturnTime          time.Time `gorm:"type:datetime;default:null;comment:'我方归还时间'" json:"return_time"`
	HisReceiveTime      time.Time `gorm:"type:datetime;default:null;comment:'对方收货时间'" json:"his_receive_time"`
	InspectCompleteTime time.Time `gorm:"type:datetime;default:null;comment:'检查完毕时间'" json:"inspect_complete_time"`
	AllSolveTime        time.Time `gorm:"type:datetime;default:null;comment:'双方解决时间'" json:"all_solve_time"`
	CompleteTime        time.Time `gorm:"type:datetime;default:null;comment:'交易完成时间'" json:"complete_time"`
	ProductPrice        float64   `gorm:"type:decimal(10,2);not null;comment:'商品价格'" json:"product_price"`
	UseDays             int       `gorm:"type:int;not null;comment:'使用天数'" json:"use_days"`
	ProductQuantity     int       `gorm:"type:int;not null;comment:'商品数量'" json:"product_quantity"`
	Freight             float64   `gorm:"type:decimal(10,2);default:null;comment:'运费'" json:"freight"`
	ActualPayment       float64   `gorm:"type:decimal(10,2);not null;comment:'实付金额'" json:"actual_payment"`
	PaymentType         int       `gorm:"type:int;not null;comment:'付款类型 0支付宝 1微信'" json:"payment_type"`
	UserID              int       `gorm:"type:int;not null;comment:'用户ID'" json:"user_id"`
	HisID               int       `gorm:"type:int;not null;comment:'对方ID'" json:"his_id"`
	MyAddressID         int       `gorm:"type:int;not null;comment:'我方地址ID'" json:"my_address_id"`
	HisAddressID        int       `gorm:"type:int;not null;comment:'对方地址ID'" json:"his_address_id"`
	ProductID           int       `gorm:"type:int;not null;comment:'商品ID'" json:"product_id"`
}

type OrderDetail struct {
	Model
	Identifier          string    `json:"identifier"`
	Status              int       `json:"status"`
	PayTime             time.Time `json:"pay_time"`
	HisDeliveryTime     time.Time `json:"his_delivery_time"`
	MyReceiveTime       time.Time `json:"my_receive_time"`
	ReturnTime          time.Time `json:"return_time"`
	HisReceiveTime      time.Time `json:"his_receive_time"`
	InspectCompleteTime time.Time `json:"inspect_complete_time"`
	AllSolveTime        time.Time `json:"all_solve_time"`
	CompleteTime        time.Time `json:"complete_time"`
	ProductPrice        float64   `json:"product_price"`
	UseDays             int       `json:"use_days"`
	ProductQuantity     int       `json:"product_quantity"`
	Freight             float64   `json:"freight"`
	ActualPayment       float64   `json:"actual_payment"`
	PaymentType         int       `json:"payment_type"`
	UserID              int       `json:"user_id"`
	HisID               int       `json:"his_id"`
	MyAddressID         int       `json:"my_address_id"`
	MyAddressName       string    `json:"my_address_name"`
	MyAddressPhone      string    `json:"my_address_phone"`
	MyAddressProvince   string    `json:"my_address_province"`
	MyAddressCity       string    `json:"my_address_city"`
	MyAddressDistrict   string    `json:"my_address_district"`
	MyAddressDetail     string    `json:"my_address_detail"`
	HisAddressID        int       `json:"his_address_id"`
	HisAddressName      string    `json:"his_address_name"`
	HisAddressPhone     string    `json:"his_address_phone"`
	HisAddressProvince  string    `json:"his_address_province"`
	HisAddressCity      string    `json:"his_address_city"`
	HisAddressDistrict  string    `json:"his_address_district"`
	HisAddressDetail    string    `json:"his_address_detail"`
	ProductID           int       `json:"product_id"`
	ProductName         string    `json:"product_name"`
	ProductImage        string    `json:"product_image"`
}

func AddOrder(order Order) int {
	database.DB.Create(&order)
	return order.ID
}

func GetOrder(orderID int) (orderDetail OrderDetail, err error) {
	err = database.DB.Table("orders").
		Select("orders.id, orders.created_at, orders.updated_at, orders.identifier, orders.status, orders.pay_time, orders.his_delivery_time, orders.my_receive_time, orders.return_time, orders.his_receive_time, orders.inspect_complete_time, orders.all_solve_time, orders.complete_time, orders.product_price, orders.use_days, orders.product_quantity, orders.freight, orders.actual_payment, orders.payment_type, orders.user_id, orders.his_id, orders.my_address_id, my_address.name as my_address_name, my_address.phone as my_address_phone, my_address.province as my_address_province, my_address.city as my_address_city, my_address.district as my_address_district, my_address.detail as my_address_detail, orders.his_address_id, his_address.name as his_address_name, his_address.phone as his_address_phone, his_address.province as his_address_province, his_address.city as his_address_city, his_address.district as his_address_district, his_address.detail as his_address_detail, orders.product_id, products.name as product_name, products.main_image as product_image").
		Joins("LEFT JOIN addresses as my_address ON orders.my_address_id = my_address.id AND my_address.deleted_at IS NULL").
		Joins("LEFT JOIN addresses as his_address ON orders.his_address_id = his_address.id AND his_address.deleted_at IS NULL").
		Joins("LEFT JOIN products ON orders.product_id = products.id AND products.deleted_at IS NULL").
		Where("orders.id = ? AND orders.deleted_at IS NULL", orderID).
		First(&orderDetail).Error
	return
}

func GetOrderByID(orderID int) (order Order) {
	database.DB.Where("id = ?", orderID).First(&order)
	return
}
func CancelOrder(orderID int) {
	database.DB.Model(&Order{}).Where("id = ?", orderID).Update("status", 0)
}
