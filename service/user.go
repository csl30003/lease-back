package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lease/config"
	"lease/dto"
	"lease/middleware"
	"lease/model"
	"lease/response"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"
)

func Login(c *gin.Context) {
	var loginReq dto.LoginReq
	var ok bool

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	//  校验用户名和密码
	var user model.User
	if user, ok = model.GetUserByNameAndPassword(loginReq.Name, loginReq.Password); !ok {
		response.Failed(c, "登录失败")
		return
	}
	//  过期时间
	expirationTime := time.Now().Add(12 * time.Hour)
	//  创建JWT声明
	claims := &middleware.Claims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//  使用用于签名的算法和令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtKey := []byte(config.Cfg.Section("JWT").Key("secret_key").String())
	//  创建JWT字符串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		response.Failed(c, "内部服务器错误")
		return
	}
	//  将客户端cookie token设置成JWT
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	var loginResp dto.LoginResp
	loginResp.Token = tokenString
	loginResp.ID = user.ID
	loginResp.Name = user.Name
	loginResp.Avatar = user.Avatar
	loginResp.Gender = user.Gender

	response.Success(c, "登录成功", loginResp)
}

func Register(c *gin.Context) {
	var registerReq dto.RegisterReq
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	if registerReq.Name == "" || registerReq.Password == "" {
		response.Failed(c, "至少需要填写用户名和密码")
		return
	}

	if ok := model.ExistUserByName(registerReq.Name); ok {
		response.Failed(c, "用户名已存在")
		return
	}

	var user model.User
	user.Name = registerReq.Name
	user.Password = registerReq.Password
	model.AddUser(&user)
	response.Success(c, "注册成功", nil)
}

func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	response.Success(c, "退出登录成功", nil)
}

func GetUserInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	user := model.GetUserByID(userId)

	response.Success(c, "获取用户信息成功", user)
}

func UpdateUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	user.ID = userId
	model.UpdateUser(&user)

	response.Success(c, "更新用户信息成功", nil)
}

func Upload(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	file, err := c.FormFile("avatar")

	if err != nil {
		response.Failed(c, "上传失败")
		return
	}

	// 构建文件名
	fileExt := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)

	// 保存上传的文件
	err = c.SaveUploadedFile(file, "D:\\HBuilderProjects\\lease\\static\\images\\avatar\\"+fileName)
	if err != nil {
		response.Failed(c, "上传失败")
		return
	}

	// 构建文件访问路径
	path := fmt.Sprintf("/static/images/avatar/%s", fileName)

	// 先删除原来的头像
	oldPath := model.GetUserByID(userId).Avatar
	if oldPath != "" {
		//删除图片
		err := os.Remove("D:\\HBuilderProjects\\lease" + oldPath)
		if err != nil {
			response.Failed(c, "删除原头像失败")
			return
		}
	}

	//保存到数据库
	model.UpdateUserAvatar(userId, path)

	response.Success(c, "上传成功", path)
}

func GetOtherUserInfo(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}
	user := model.GetUserByID(idInt)
	response.Success(c, "获取用户信息成功", user)
}
