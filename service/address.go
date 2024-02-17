package service

import (
	"github.com/gin-gonic/gin"
	"lease/model"
	"lease/response"
	"reflect"
	"strconv"
)

func GetAddress(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	addressList := model.GetAddressListByUserID(userId)

	response.Success(c, "获取成功", addressList)
}

func GetAddressByID(c *gin.Context) {
	addressId, _ := strconv.Atoi(c.Param("id"))
	address := model.GetAddressByID(addressId)

	response.Success(c, "获取成功", address)
}

func SetDefaultAddress(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	addressId, _ := strconv.Atoi(c.Param("id"))
	model.SetDefaultAddress(addressId, userId)

	response.Success(c, "设置成功", nil)
}

func AddAddress(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	var address model.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	address.UserID = userId
	model.AddAddress(address)

	response.Success(c, "添加成功", nil)
}

func UpdateAddress(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	addressId, _ := strconv.Atoi(c.Param("id"))
	var address model.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	address.UserID = userId
	address.ID = addressId
	model.UpdateAddress(address)

	response.Success(c, "更新成功", nil)
}

func DeleteAddress(c *gin.Context) {
	addressId, _ := strconv.Atoi(c.Param("id"))
	model.DeleteAddress(addressId)

	response.Success(c, "删除成功", nil)
}
