package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("HelloJwt")

const (
	SUCCESS = iota
	InvalidParams
	ErrorAuthCheckTokenFail
	ErrorAuthCheckTokenTimeout
)

type Claims struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userName, password string, role int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		userName,
		password,
		role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = SUCCESS
		token := c.GetHeader("token")
		// token := c.Query("token")
		if token == "" {
			code = InvalidParams
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = ErrorAuthCheckTokenTimeout
			}
		}

		if code != SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "check token failed",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
