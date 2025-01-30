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
	marshResponse, err := json.Marshal(items[1])
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
