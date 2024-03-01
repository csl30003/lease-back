package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lease/dto"
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

func UpdateProductStatus(c *gin.Context) {
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

func GetProductList(c *gin.Context) {
	//从get方法获取数据
	current := c.Query("current")
	size := c.Query("size")
	name := c.Query("name")
	categoryID := c.Query("category_id")
	sort := c.Query("sort")
	order := c.Query("order")

	// 默认判断
	if current == "" {
		current = "1"
	}
	if size == "" {
		size = "10"
	}
	if sort == "" {
		sort = "0"
	}
	if order == "" {
		order = "asc"
	}

	//为0表示时间排序，为1表示使用时间排序，为2表示使用价格排序
	if sort == "0" {
		sort = "id"
	} else if sort == "1" {
		sort = "fineness"
	} else {
		sort = "price"
	}

	// 将current和size转换成int类型
	currentInt, err := strconv.Atoi(current)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	productList := model.GetProductList(currentInt, sizeInt, name, categoryID, sort, order)
	total := model.GetProductListTotal(name, categoryID, sort, order)

	var pages int64
	if total%int64(sizeInt) == 0 {
		pages = total / int64(sizeInt)
	} else {
		pages = total/int64(sizeInt) + 1
	}

	response.Success(c, "获取成功", gin.H{
		"records": productList,
		"pages":   pages,
	})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	product := model.GetProduct(productId)

	// 通过product.CategoryID获取分类信息
	category := model.GetCategoryByID(product.CategoryID)

	// 通过product.AddressID获取地址信息
	address := model.GetAddressByID(product.AddressID)

	// 通过product.UserID获取用户信息
	user := model.GetUserByID(product.UserID)

	// 通过product.ID获取商品图片列表
	imageList := model.GetProductImageList(product.ID)

	getProductResp := dto.GetProductResp{
		ID:              product.ID,
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
		DeletedAt:       product.DeletedAt,
		Name:            product.Name,
		Price:           product.Price,
		Detail:          product.Detail,
		MainImage:       product.MainImage,
		Stock:           product.Stock,
		Delivery:        product.Delivery,
		Freight:         product.Freight,
		Fineness:        product.Fineness,
		UsedYears:       product.UsedYears,
		Status:          product.Status,
		CategoryID:      category.ID,
		CategoryName:    category.Name,
		AddressID:       address.ID,
		AddressName:     address.Name,
		AddressPhone:    address.Phone,
		AddressProvince: address.Province,
		AddressCity:     address.City,
		AddressDistrict: address.District,
		AddressDetail:   address.Detail,
		UserID:          user.ID,
		UserName:        user.Name,
		UserAvatar:      user.Avatar,
		ImageList:       imageList,
	}

	response.Success(c, "获取成功", getProductResp)
}

func GetCarousel(c *gin.Context) {
	carouselList := model.GetCarousel()
	response.Success(c, "获取成功", carouselList)
}
