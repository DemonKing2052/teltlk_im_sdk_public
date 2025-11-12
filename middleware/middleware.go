package middleware

import (
	"ImSdk/common"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Token
		tokenString := c.GetHeader("Authorization")

		// 检查Token是否存在
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未提供Token"})
			c.Abort()
			return
		}

		// Token通常以"Bearer "开头，去掉前缀
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "无效的Token格式"})
			c.Abort()
			return
		}

		// 解析Token
		tokenString = parts[1]
		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return common.JwtKey, nil
		})

		fmt.Print(claims.Subject)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "无效的Token"})
			c.Abort()
			return
		}

		// 将解析的Token中的用户信息存储在上下文中，以便后续处理函数使用
		c.Set("userID", claims.Subject)

		// 继续处理请求
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, applet-name, token")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// GinCors gin跨域中间件
func GinCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method //请求方法
		c.Header("Access-Control-Allow-Origin", "*")
		//由于跨域不知道请求头多了什么东西，暂时使用*
		c.Header("Access-Control-Allow-Headers", "*")
		//c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
