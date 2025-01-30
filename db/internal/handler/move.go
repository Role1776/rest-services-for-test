package handler

import (
	"app"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateDB(c *gin.Context) {
	var userTable app.UserTable
	if err := c.BindJSON(&userTable); err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "details": "Failed to bind JSON to UserTable struct"})
		return
	}

	fmt.Printf("Received request to create list. UserID: %d, Title: %s, Description: %s\n", userTable.UserId, userTable.Title, userTable.Description)

	id, err := h.repos.CreateList(userTable)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "details": "Failed to create list in database"})
		return
	}

	c.JSON(200, id)
}

func (h *handler) GetDB(c *gin.Context) {
	var id int
	if err := c.BindJSON(&id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tables, err := h.repos.GetListsUser(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tables)
}
