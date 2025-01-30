package handler

import (
	"app"

	"github.com/gin-gonic/gin"
)

func (h *handler) SignUpDB(c *gin.Context) {
	var user app.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	id, err := h.repos.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, id)
}

func (h *handler) SignInDB(c *gin.Context) {
	var user app.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	id, err := h.repos.GetUser(user.Username, user.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}
