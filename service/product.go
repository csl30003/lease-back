package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lease/model"
	"lease/response"
	"reflect"
	"strconv"
)

func AddProduct(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		fmt.Println(err)
		response.Failed(c, "参数错误")
		return
	}
	product.UserID = userId
	productId := model.AddProduct(product)

	fmt.Println(productId)

	response.Success(c, "添加成功", productId)
}

func UpdateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	model.UpdateProductStatus(product)
	response.Success(c, "更新成功", nil)
}

func GetMyProduct(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	status, _ := strconv.Atoi(c.Param("status"))

	products := model.GetMyProduct(userId, status)
	response.Success(c, "获取成功", products)
}
