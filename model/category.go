package model

import "lease/database"

type Category struct {
	Model
	ParentID int    `gorm:"column:parent_id;type:int;not null;default:0;comment:父类别ID" json:"parent_id"`
	Name     string `gorm:"column:name;type:varchar(255);not null;comment:种类名称" json:"name"`
	Sort     int    `gorm:"column:sort;type:int;not null;default:0;comment:排序" json:"sort"`
	Status   int    `gorm:"column:status;type:tinyint(1);default:0;comment:状态" json:"status"`
}

func GetCategoryByParentID(parentId int) (categoryList []Category) {
	database.DB.Where("parent_id = ?", parentId).Find(&categoryList)
	return
}
