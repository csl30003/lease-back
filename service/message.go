package service

import (
	"github.com/gin-gonic/gin"
	"lease/model"
	"lease/response"
	"reflect"
	"strconv"
)

func SendMessage(c *gin.Context) {
	var message model.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	if message.FromID == 0 || message.ToID == 0 || message.Content == "" {
		response.Failed(c, "参数错误")
		return
	}

	model.AddMessage(message)

	response.Success(c, "发送成功", nil)
}

func GetChatUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	messageUserList := model.GetChatUser(userId)

	response.Success(c, "获取成功", messageUserList)
}

func GetMessage(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	toId := c.Param("toId")
	toIdInt, err := strconv.Atoi(toId)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	messageList := model.GetMessage(userId, toIdInt)

	response.Success(c, "获取成功", messageList)
}
