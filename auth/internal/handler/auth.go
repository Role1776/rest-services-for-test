package handler

import (
	"app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) SignUp(c *gin.Context) {
	var user app.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	user, err := h.services.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	resp, err := http.Post("http://localhost:8000/auth/sign-up", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.JSON(400, gin.H{"error": err})
		return
	}

	var id int
	if err := json.NewDecoder(resp.Body).Decode(&id); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, id)
}

func (h *handler) SignIn(c *gin.Context) {
	var user app.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	user.Password = h.services.GeneratePasswordHash(user.Password)
	userBytes, err := json.Marshal(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	resp, err := http.Post("http://localhost:8000/auth/sign-in", "application/json", bytes.NewBuffer(userBytes))
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.JSON(400, gin.H{"error": err})
		return
	}

	var response struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	jwt, err := h.services.GenerateJWT(response.ID)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, jwt)

}

func (h *handler) Token(c *gin.Context) {
	var token string
	if err := c.BindJSON(&token); err != nil {
		c.JSON(400, gin.H{"error of binding": err})
		return
	}

	id, err := h.services.ParseToken(token)
	if err != nil {
		c.JSON(400, gin.H{"error of parsing": err})
		return
	}

	fmt.Printf("User ID: %d\n", id)

	c.JSON(200, id)
}
