package main

import (
	"bytes"
	"crypto/sha256"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")

		if err != nil {
			c.Data(http.StatusUnauthorized, RETURNTYPE, []byte("token error"))
			c.Abort()
			return
		}
	}
}

func AuthUser(username string, password string) bool {
	sha := sha256.New()
	sha.Write([]byte(password))
	hash := sha.Sum(nil)

	if passwordHash, ok := GetPasswordHash(username); ok {
		return bytes.Equal(hash, []byte(passwordHash))
	}
	return false
}
