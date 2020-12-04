package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const strToken = "mysecret"

func generateToken(str string) (string, error) {
	fmt.Print()
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte(str))
	if err != nil {
		return "", err
	}
	return token, err
}

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	bearerToken := strings.Split(tokenString, " ")
	if len(bearerToken) == 2 {
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(strToken), nil
		})

		if token != nil && err == nil {
			fmt.Println("token verified")
		} else {
			result := gin.H{
				"message": "not authorized",
				"error":   err.Error(),
			}
			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
		}
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   "An authorization header is required",
		}

		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}

}
