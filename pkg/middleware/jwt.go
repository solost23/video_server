package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header中获取
		tknStr := c.Request.Header.Get("Auth")
		if tknStr == "" {
			c.JSON(http.StatusBadRequest, "Auth err")
			c.Abort()
			return
		}
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWTKEY), nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		if !tkn.Valid {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
		return
	}
}

const JWTKEY = "my_secret_key"

type Claims struct {
	UserName string
	Role     string
	jwt.StandardClaims
}

func CreateToken(userName, role string) (string, error) {
	// 后续修改为配置
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		UserName: userName,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTKEY))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return tokenString, nil
}
