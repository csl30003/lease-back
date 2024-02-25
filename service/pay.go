package service

import (
	"github.com/gin-gonic/gin"
	"lease/config"
	"lease/model"
	"lease/response"
	"lease/util"
	"strconv"
)

const (
	kAppID               = "9021000134675991"
	kServerDomain        = "http://e8967k.natappfree.cc"
	AppPublicCertPath    = "config/dev/cert/appPublicCert.crt"    // app公钥证书路径
	AliPayRootCertPath   = "config/dev/cert/alipayRootCert.crt"   // alipay根证书路径
	AliPayPublicCertPath = "config/dev/cert/alipayPublicCert.crt" // alipay公钥证书路径
	NotifyURL            = kServerDomain + "/notify"
	ReturnURL            = kServerDomain + "/callback"
	IsProduction         = false
)

var AliPayClient *util.AliPayClient

func init() {
	kPrivateKey := config.Cfg.Section("AliPay").Key("k_private_key").String()
	AliPayClient = util.InitClient(util.Config{
		KAppID:               kAppID,
		KPrivateKey:          kPrivateKey,
		IsProduction:         IsProduction,
		AppPublicCertPath:    AppPublicCertPath,
		AliPayRootCertPath:   AliPayRootCertPath,
		AliPayPublicCertPath: AliPayPublicCertPath,
		NotifyURL:            NotifyURL,
		ReturnURL:            ReturnURL,
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

	// 商家钱包增加金额
	order := model.GetOrderByID(orderIDInt)
	user := model.GetUserByID(order.UserID)
	model.UpdateUserWallet(user.ID, user.Wallet+order.ActualPayment)

	response.Success(c, "支付成功", nil)
}
