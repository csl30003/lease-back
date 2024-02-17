package model

import "lease/database"

type User struct {
	Model
	Name     string `gorm:"column:name;type:varchar(255);not null;comment:昵称;unique" json:"name"`
	Password string `gorm:"column:password;type:varchar(255);not null;comment:密码" json:"password"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);default:null;comment:头像" json:"avatar"`
	Gender   int    `gorm:"column:gender;type:tinyint(1);default:0;comment:性别 0 男 1 女" json:"gender"`
	Country  string `gorm:"column:country;type:varchar(100);default:null;comment:所在国家" json:"country"`
	Province string `gorm:"column:province;type:varchar(100);default:null;comment:省份" json:"province"`
	City     string `gorm:"column:city;type:varchar(100);default:null;comment:城市" json:"city"`
	Phone    string `gorm:"column:phone;type:varchar(255);default:null;comment:电话" json:"phone"`
}

func GetUserByNameAndPassword(name, password string) (user User, ok bool) {
	if err := database.DB.Where("name = ? and password = ?", name, password).First(&user).Error; err != nil {
		return user, false
	}
	return user, true
}

func ExistUserByName(name string) bool {
	var user User
	if err := database.DB.Where("name = ?", name).First(&user).Error; err != nil {
		return false
	}
	return true
}

func AddUser(user *User) {
	database.DB.Create(user)
}
