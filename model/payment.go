package model

import (
	"lease/database"
	"time"
)

type Payment struct {
	Model
	Type    int     `gorm:"type:tinyint(1);not null;comment:'类型 1表示收入 2表示提现 3表示提现中'" json:"type"`
	Money   float64 `gorm:"type:decimal(10,2);not null;comment:'金额'" json:"money"`
	UserID  int     `gorm:"not null;comment:'用户ID'" json:"user_id"`
	OrderID int     `gorm:"default:0;not null;comment:'订单ID（仅在收入时有效）'" json:"order_id"`
}

func AddPayment(payment Payment) {
	database.DB.Create(&payment)
}

type PaymentDetail struct {
	ID              int       `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	Type            int       `json:"type"`
	Money           float64   `json:"money"`
	UserID          int       `json:"user_id"`
	OrderID         int       `json:"order_id"`
	OrderIdentifier string    `json:"order_identifier"`
}

func GetPaymentByUserID(userId int) (paymentList []PaymentDetail) {
	database.DB.Table("payments").
		Select("payments.id, payments.created_at, payments.type, payments.money, payments.user_id, payments.order_id, orders.identifier as order_identifier").
		Joins("left join orders on payments.order_id = orders.id AND orders.deleted_at IS NULL").
		Where("payments.user_id = ? AND payments.type != 3 AND payments.deleted_at IS NULL", userId).
		Order("created_at desc").
		Find(&paymentList)
	return
}
