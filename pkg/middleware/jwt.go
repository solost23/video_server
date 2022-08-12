package middleware

import (
	"fmt"
	"net/http"
	"time"
	"video_server/config"
	"video_server/pkg/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header中获取
		tknStr := c.Request.Header.Get("token")
		if tknStr == "" {
			c.JSON(http.StatusBadRequest, "token err")
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
		c.Set("user", claims.User)
		c.Next()
		return
	}
}

var JWTKEY = config.JwtKey

type Claims struct {
	User *models.User
	jwt.StandardClaims
}

func CreateToken(user *models.User) (string, error) {
	// 后续修改为配置
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		User: user,
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
