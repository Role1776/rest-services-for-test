package handler

import (
	"app"
	"app/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) SignUp(c *gin.Context) {
	var user app.User
	if err := c.BindJSON(&user); err != nil {
		logger.Log().Error("Error binding JSON: %v", err)
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid request body: %v", err)})
		return
	}
	logger.Log().Info("User data bound successfully")

	us, err := json.Marshal(user)
	if err != nil {
		logger.Log().Error("Error marshalling user: %v", err)
		c.JSON(400, gin.H{"error": fmt.Sprintf("Failed to process user data: %v", err)})
		return
	}
	logger.Log().Info("User data marshalled successfully")

	firstResp, err := http.Post("http://localhost:8001/auth/sign-up", "application/json", bytes.NewBuffer(us))
	if err != nil {
		logger.Log().Error("Error making POST request: %v", err)
		c.JSON(400, gin.H{"error": fmt.Sprintf("Failed to communicate with authentication service: %v", err)})
		return
	}
	logger.Log().Info("POST request to authentication service completed")

	defer firstResp.Body.Close()
	if firstResp.StatusCode != 200 {
		logger.Log().Error("Received non-200 status code: %d", firstResp.StatusCode)
		c.JSON(400, gin.H{"error": fmt.Sprintf("Authentication service returned status code: %d", firstResp.StatusCode)})
		return
	}
	logger.Log().Info("Received 200 status code from authentication service")

	var id int
	if err := json.NewDecoder(firstResp.Body).Decode(&id); err != nil {
		logger.Log().Error("Error decoding response body: %v", err)
		c.JSON(400, gin.H{"error": fmt.Sprintf("Failed to process response from authentication service: %v", err)})
		return
	}

	logger.Log().Info("User signed up successfully with ID: %d", id)
	c.JSON(200, gin.H{"id": id})
}

func (h *handler) SignIn(c *gin.Context) {
	var user app.User
	if err := c.BindJSON(&user); err != nil {
		logger.Log().Error("Error binding JSON: %v", err)
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid request body: %v", err)})
		return
	}
	logger.Log().Info("User data bound successfully")

	us, err := json.Marshal(user)
	if err != nil {
		logger.Log().Error("Error marshalling user: %v", err)
		c.JSON(400, gin.H{"error": fmt.Sprintf("Failed to process user data: %v", err)})
		return
	}
	logger.Log().Info("User data marshalled successfully")

	firstResp, err := http.Post("http://localhost:8001/auth/sign-in", "application/json", bytes.NewBuffer(us))
	if err != nil {
		logger.Log().Error("Error making POST request: %v", err)
		c.JSON(400, gin.H{"error": fmt.Sprintf("Failed to communicate with authentication service: %v", err)})
		return
	}
	logger.Log().Info("POST request to authentication service completed")

	defer firstResp.Body.Close()
	if firstResp.StatusCode != 200 {
		logger.Log().Error("Received non-200 status code: %d", firstResp.StatusCode)
		c.JSON(400, gin.H{"error": fmt.Sprintf("Authentication failed with status code: %d", firstResp.StatusCode)})
		return
	}
	logger.Log().Info("Received 200 status code from authentication service")

	var token string
	if err := json.NewDecoder(firstResp.Body).Decode(&token); err != nil {
		logger.Log().Error("Error decoding response body: %v", err)
		c.JSON(400, gin.H{"error": fmt.Sprintf("Failed to process response from authentication service: %v", err)})
		return
	}

	logger.Log().Info("User signed in successfully")
	c.JSON(200, gin.H{"token": token})
}
