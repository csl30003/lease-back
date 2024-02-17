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
	"path/filepath"
	"reflect"
	"strconv"
	"time"
)

func HelloWorld(c *gin.Context) {
	// 获取参数
	man := struct {
		Name string `json:"name"`
		Sex  int    `json:"sex"`
	}{}
	if err := c.ShouldBindJSON(&man); err != nil {
		c.String(400, "请求参数错误")
		return
	}
	fmt.Println(man.Name, man.Sex)
	// sex转字符串
	strSex := strconv.Itoa(man.Sex)
	//	返回一个字符串
	c.String(200, "Hello %s %s", man.Name, strSex)
}

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
	expirationTime := time.Now().Add(60 * 60 * time.Second) //12 * time.Hour
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

// 还没写完
func Upload(c *gin.Context) {
	fmt.Println("11111")
	file, err := c.FormFile("avatar")

	if err != nil {
		response.Failed(c, "上传失败")
		return
	}
	fmt.Println(file.Filename)

	// 构建文件名
	fileExt := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)
	// 保存上传的文件
	fmt.Println(fileName)
	//err = c.SaveUploadedFile(file, "./uploads/"+fileName)
	//if err != nil {
	//	response.Failed(c, "上传失败")
	//	return
	//}

	// 构建文件访问路径
	path := fmt.Sprintf("http://localhost:8080/uploads/%s", fileName)
	response.Success(c, "上传成功", path)
}
