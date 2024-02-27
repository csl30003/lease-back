package service

import (
	"github.com/gin-gonic/gin"
	"lease/common"
	"lease/config"
	"lease/database"
	"lease/dto"
	"lease/model"
	"lease/response"
	"lease/util"
	"reflect"
	"strconv"
)

const (
	kAppID               = "9021000134675991"
	AppPublicCertPath    = "config/dev/cert/appPublicCert.crt"    // app公钥证书路径
	AliPayRootCertPath   = "config/dev/cert/alipayRootCert.crt"   // alipay根证书路径
	AliPayPublicCertPath = "config/dev/cert/alipayPublicCert.crt" // alipay公钥证书路径
	IsProduction         = false
)

var AliPayClient *util.AliPayClient

func init() {
	kServerDomain := config.Cfg.Section("AliPay").Key("k_server_domain").String()
	notifyURL := kServerDomain + "/notify"
	returnURL := kServerDomain + "/callback"
	kPrivateKey := config.Cfg.Section("AliPay").Key("k_private_key").String()
	AliPayClient = util.InitClient(util.Config{
		KAppID:               kAppID,
		KPrivateKey:          kPrivateKey,
		IsProduction:         IsProduction,
		AppPublicCertPath:    AppPublicCertPath,
		AliPayRootCertPath:   AliPayRootCertPath,
		AliPayPublicCertPath: AliPayPublicCertPath,
		NotifyURL:            notifyURL,
		ReturnURL:            returnURL,
	})
}

// PayUrl 重定向到支付宝二维码
func PayUrl(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	// 获取订单价格
	order := model.GetOrderByID(idInt)
	url, err := AliPayClient.Pay(util.Order{
		ID:          strconv.Itoa(order.ID),
		Subject:     order.Identifier,
		TotalAmount: float32(order.ActualPayment),
		Code:        util.LaptopWebPay,
	})

	if err != nil {
		response.Failed(c, "系统错误")
		return
	}

	response.Success(c, "进入支付页面", url)
}

// Callback 支付后页面的重定向界面
func Callback(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		response.Failed(c, "支付失败")
		return
	}
	response.Success(c, "支付成功", orderID)
}

// Notify 支付成功后支付宝异步通知
func Notify(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		response.Failed(c, "支付失败")
		return
	}
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		response.Failed(c, "支付失败，订单id错误")
		return
	}

	// 改变订单状态
	model.UpdateOrderStatus(orderIDInt, 2)

	// 生成订单pay_time
	model.UpdateOrderPayTime(orderIDInt)

	// 在收支表中增加一条记录
	order := model.GetOrderByID(orderIDInt)
	model.AddPayment(model.Payment{
		Type:    1,
		Money:   order.ActualPayment,
		UserID:  order.HisID,
		OrderID: orderIDInt,
	})

	// 商家钱包增加金额
	user := model.GetUserByID(order.HisID)
	model.UpdateUserWallet(user.ID, user.Wallet+order.ActualPayment)

	response.Success(c, "支付成功", nil)
}

func GetPayment(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	paymentList := model.GetPaymentByUserID(userId)

	getPaymentResp := make([]dto.Payment, 0)
	for _, payment := range paymentList {
		getPaymentResp = append(getPaymentResp, dto.Payment{
			ID:              payment.ID,
			CreatedAt:       payment.CreatedAt.Format("2006-01-02 15:04:05"),
			Type:            payment.Type,
			Money:           payment.Money,
			UserID:          payment.UserID,
			OrderID:         payment.OrderID,
			OrderIdentifier: payment.OrderIdentifier,
		})
	}

	response.Success(c, "获取成功", getPaymentResp)
}

func Withdraw(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	// 判断是否有提现中的记录
	val, _ := database.Cache.Get(common.RedisKeyWithdrawUserID + strconv.Itoa(userId)).Result()
	if val == "" {
		// 获取用户信息
		user := model.GetUserByID(userId)
		if user.Wallet == 0 {
			response.Failed(c, "钱包为空")
			return
		}

		// 提现
		model.AddPayment(model.Payment{
			Type:   3,
			Money:  user.Wallet,
			UserID: userId,
		})

		// 在redis中增加提现中的记录
		database.Cache.Set(common.RedisKeyWithdrawUserID+strconv.Itoa(userId), 1, 0)

		response.Success(c, "提现申请成功", nil)
	} else {
		response.Failed(c, "已申请提现，请勿重复提现")
	}
}
