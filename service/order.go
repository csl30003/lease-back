package service

import (
	"github.com/gin-gonic/gin"
	"lease/dto"
	"lease/model"
	"lease/response"
	"lease/util"
	"reflect"
	"strconv"
	"time"
)

func AddOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	// 减少商品库存
	product := model.GetProduct(order.ProductID)
	if product.Stock < order.ProductQuantity {
		response.Failed(c, "库存不足")
		return
	}
	product.Stock -= order.ProductQuantity
	model.UpdateProductStock(product.ID, product.Stock)

	// 生成订单号
	t := time.Now()
	order.Identifier = util.Generate(t)

	id := model.AddOrder(order)

	response.Success(c, "添加成功", id)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	order, err := model.GetOrder(idInt)
	if err != nil {
		response.Failed(c, "订单不存在")
		return
	}

	getOrderResp := dto.GetOrderResp{
		ID:                  order.ID,
		CreatedAt:           order.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:           order.UpdatedAt.Format("2006-01-02 15:04:05"),
		Identifier:          order.Identifier,
		Status:              order.Status,
		PayTime:             order.PayTime.Format("2006-01-02 15:04:05"),
		HisDeliveryTime:     order.HisDeliveryTime.Format("2006-01-02 15:04:05"),
		MyReceiveTime:       order.MyReceiveTime.Format("2006-01-02 15:04:05"),
		ReturnTime:          order.ReturnTime.Format("2006-01-02 15:04:05"),
		HisReceiveTime:      order.HisReceiveTime.Format("2006-01-02 15:04:05"),
		InspectCompleteTime: order.InspectCompleteTime.Format("2006-01-02 15:04:05"),
		AllSolveTime:        order.AllSolveTime.Format("2006-01-02 15:04:05"),
		CompleteTime:        order.CompleteTime.Format("2006-01-02 15:04:05"),
		ProductPrice:        order.ProductPrice,
		UseDays:             order.UseDays,
		ProductQuantity:     order.ProductQuantity,
		Freight:             order.Freight,
		ActualPayment:       order.ActualPayment,
		PaymentType:         order.PaymentType,
		UserID:              order.UserID,
		HisID:               order.HisID,
		MyAddressID:         order.MyAddressID,
		MyAddressName:       order.MyAddressName,
		MyAddressPhone:      order.MyAddressPhone,
		MyAddressProvince:   order.MyAddressProvince,
		MyAddressCity:       order.MyAddressCity,
		MyAddressDistrict:   order.MyAddressDistrict,
		MyAddressDetail:     order.MyAddressDetail,
		HisAddressID:        order.HisAddressID,
		HisAddressName:      order.HisAddressName,
		HisAddressPhone:     order.HisAddressPhone,
		HisAddressProvince:  order.HisAddressProvince,
		HisAddressCity:      order.HisAddressCity,
		HisAddressDistrict:  order.HisAddressDistrict,
		HisAddressDetail:    order.HisAddressDetail,
		ProductID:           order.ProductID,
		ProductName:         order.ProductName,
		ProductImage:        order.ProductImage,
	}

	response.Success(c, "获取成功", getOrderResp)
}

func CancelOrder(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	order := model.GetOrderByID(idInt)
	if order.UserID != userId {
		response.Failed(c, "无权限取消订单")
		return
	}

	// 还库存
	product := model.GetProduct(order.ProductID)
	product.Stock += order.ProductQuantity
	model.UpdateProductStock(product.ID, product.Stock)

	model.UpdateOrderStatus(idInt, 0)

	response.Success(c, "取消成功", nil)
}

func GetMyOrder(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	//从get方法获取数据
	current := c.Query("current")
	size := c.Query("size")
	status := c.Query("status")

	// 默认判断
	if current == "" {
		current = "1"
	}
	if size == "" {
		size = "10"
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

	var orderList []model.OrderDetail
	var total int64
	// 如果status在0到8之间，查询指定状态订单，否则查询全部订单
	if status == "0" || status == "1" || status == "2" || status == "3" || status == "4" || status == "5" || status == "6" || status == "7" || status == "8" {
		statusInt, err := strconv.Atoi(status)
		if err != nil {
			response.Failed(c, "参数错误")
			return
		}
		// 查询指定状态订单
		orderList = model.GetMyPartialOrder(userId, currentInt, sizeInt, statusInt)
		total = model.GetMyPartialOrderTotal(userId, statusInt)
	} else {
		// 查询全部订单
		orderList = model.GetMyAllOrder(userId, currentInt, sizeInt)
		total = model.GetMyAllOrderTotal(userId)
	}

	var pages int64
	if total%int64(sizeInt) == 0 {
		pages = total / int64(sizeInt)
	} else {
		pages = total/int64(sizeInt) + 1
	}

	response.Success(c, "获取成功", gin.H{
		"records": orderList,
		"pages":   pages,
	})
}

// IReceiveOrder 我收货
func IReceiveOrder(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	order := model.GetOrderByID(idInt)
	if order.UserID != userId {
		response.Failed(c, "无权限操作")
		return
	}

	model.UpdateOrderStatus(idInt, 4)

	model.UpdateOrderMyReceiveTime(idInt)

	response.Success(c, "操作成功", nil)
}

// IReturnOrder 我归还
func IReturnOrder(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	order := model.GetOrderByID(idInt)
	if order.UserID != userId {
		response.Failed(c, "无权限操作")
		return
	}

	model.UpdateOrderStatus(idInt, 5)

	model.UpdateOrderReturnTime(idInt)

	response.Success(c, "操作成功", nil)
}

// SolveOrder 双方确认解决 没写完
func SolveOrder(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	order := model.GetOrderByID(idInt)
	if order.UserID != userId && order.HisID != userId {
		response.Failed(c, "无权限操作")
		return
	}

	// 用redis判断是否双方都确认解决了

	response.Success(c, "操作成功", nil)
}
