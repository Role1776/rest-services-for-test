package handler

import (
	"app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Create(c *gin.Context) {
	id, ok := c.Get(userIdentity)
	if !ok {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	var userTable app.UserTable
	if err := c.BindJSON(&userTable); err != nil {
		c.JSON(400, gin.H{"error of binding gateway": err})
		return
	}

	fmt.Printf("Creating table for user ID: %v\n", id)

	userTable.UserId = id.(int)
	jsonUserTable, err := json.Marshal(userTable)
	if err != nil {
		c.JSON(400, gin.H{"error of marshalling gateway": err})
		return
	}

	resp, err := http.Post("http://localhost:8000/api/move/create", "application/json", bytes.NewBuffer(jsonUserTable))
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Failed to send POST request: %v", err)})
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid response body. Status code: %d", resp.StatusCode)})
		return
	}

	var idUser int
	if err := json.NewDecoder(resp.Body).Decode(&idUser); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Failed to decode response ID: %v", err)})
		return
	}

	c.JSON(200, gin.H{"id": idUser})

}

func (h *handler) Get(c *gin.Context) {
	id, ok := c.Get(userIdentity)
	if !ok {
		c.JSON(400, gin.H{"error": "User not found. Unable to retrieve user identity from context."})
		return
	}
	jsonId, err := json.Marshal(id.(int))
	if err != nil {
		c.JSON(400, gin.H{"error of marshalling": err})
		return
	}

	resp, err := http.Post("http://localhost:8000/api/move", "application/json", bytes.NewBuffer(jsonId))
	if err != nil {
		c.JSON(400, gin.H{"error of posting": err})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.JSON(400, gin.H{"error of status code": resp.StatusCode})
		return
	}

	var lists []app.UserTable
	if err := json.NewDecoder(resp.Body).Decode(&lists); err != nil {
		c.JSON(400, gin.H{"error of decoding": err})
		return
	}

	c.JSON(200, gin.H{"message": lists})

}

//{
//"user_id":6,
// "list_id":11,
// "title": "устроится в гугл2",
// "description": "ешкере2"
//}
