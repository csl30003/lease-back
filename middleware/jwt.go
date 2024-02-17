package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lease/config"
	"lease/response"
	"net/http"
)

//
// Claims
//  @Description: JWT返回的JSON
//
type Claims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//  获取Cookie
		ck, err := c.Request.Cookie("token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				// 没有设置Cookie
				response.Unauthorized(c, "未授权")
				c.Abort()
				return
			}
			response.Unauthorized(c, "未知错误")
			c.Abort()
			return
		}

		tokenString := ck.Value
		claims := &Claims{}

		jwtKey := []byte(config.Cfg.Section("JWT").Key("secret_key").String())

		//  解析JWT字符串并吧结果存储在claims中
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			fmt.Println(err)
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				response.Unauthorized(c, "未授权")
				c.Abort()
				return
			}
			response.Unauthorized(c, "未知错误")
			c.Abort()
			return
		}
		if !token.Valid {
			response.Unauthorized(c, "未授权")
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
