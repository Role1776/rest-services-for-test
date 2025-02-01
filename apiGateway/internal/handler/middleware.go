package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	userIdentity = "userId"
)

func (h *handler) userIdentityMiddleware(c *gin.Context) {
	var tokenString string
	if cookie, err := c.Cookie("token"); cookie != "" && err == nil {
		tokenString = cookie
	} else {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "unauthorized"})
			return
		}

		items := strings.Split(token, " ")
		if len(items) != 2 || items[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"message": "invalid auth header"})
			return
		}
		tokenString = items[1]

		c.SetCookie("token", tokenString, 18000, "/", "", false, true)
	}

	marshResponse, err := json.Marshal(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": err})
		return
	}

	resp, err := http.Post("http://localhost:8001/auth/token", "application/json", bytes.NewBuffer(marshResponse))
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": err})
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.AbortWithStatusJSON(401, gin.H{"message": err})
		return
	}
	var userId int
	json.NewDecoder(resp.Body).Decode(&userId)

	c.Set(userIdentity, userId)
	c.Next()
}
