package model

import "lease/database"

type Address struct {
	Model
	Name      string `gorm:"column:name;type:varchar(255);not null;comment:姓名" json:"name"`
	Phone     string `gorm:"column:phone;type:varchar(255);not null;comment:电话" json:"phone"`
	Province  string `gorm:"column:province;type:varchar(255);not null;comment:省份" json:"province"`
	City      string `gorm:"column:city;type:varchar(255);not null;comment:城市" json:"city"`
	District  string `gorm:"column:district;type:varchar(255);not null;comment:区县" json:"district"`
	Detail    string `gorm:"column:detail;type:varchar(255);not null;comment:详细地址" json:"detail"`
	IsDefault int    `gorm:"column:is_default;type:tinyint(1);default:0;comment:是否为默认地址" json:"is_default"`
	UserID    int    `gorm:"column:user_id;type:int;not null;comment:用户ID" json:"user_id"`
}

func GetAddressListByUserID(userId int) (addressList []Address) {
	database.DB.Where("user_id = ?", userId).Find(&addressList)
	return
}

func GetAddressByID(addressId int) (address Address) {
	database.DB.Where("id = ?", addressId).First(&address)
	return
}

func SetDefaultAddress(addressId, userId int) {
	database.DB.Model(&Address{}).Where("user_id = ?", userId).Update("is_default", 0)
	database.DB.Model(&Address{}).Where("id = ?", addressId).Update("is_default", 1)
}

func AddAddress(address Address) {
	database.DB.Create(&address)
}

func UpdateAddress(address Address) {
	database.DB.Model(&Address{}).Where("id = ?", address.ID).Updates(map[string]interface{}{"name": address.Name, "phone": address.Phone, "province": address.Province, "city": address.City, "district": address.District, "detail": address.Detail})
}

func DeleteAddress(addressId int) {
	database.DB.Delete(&Address{}, addressId)
}
