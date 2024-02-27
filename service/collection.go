package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lease/model"
	"lease/response"
	"reflect"
	"strconv"
)

func IsCollection(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	productID := c.Param("product_id")
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	if id := model.ExistCollection(productIDInt, userId); id > 0 {
		response.Success(c, "已收藏", true)
		return
	}
	response.Success(c, "未收藏", false)
}

func Collection(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	var collection model.Collection
	if err := c.ShouldBindJSON(&collection); err != nil {
		fmt.Println(err)
		response.Failed(c, "参数错误")
		return
	}

	collection.UserID = userId

	// 判断数据库是否存在该收藏
	if id := model.ExistCollection(collection.ProductID, collection.UserID); id > 0 {
		collection.ID = id
		model.DeleteCollection(collection)
		response.Success(c, "取消收藏成功", false)
		return
	}
	model.AddCollection(collection)
	response.Success(c, "收藏成功", true)

}

func GetCollection(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	collectionList := model.GetCollection(userId)

	response.Success(c, "获取收藏成功", collectionList)
}
