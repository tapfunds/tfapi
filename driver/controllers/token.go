package controllers

import (
	"fmt"
	"net/http"
	"tfdb/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindPlaidInfos ... Get all users
// GET /tokens
// Get all tokens
func FindPlaidInfos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var info []models.PlaidIntegration
	db.Find(&info)

	c.JSON(http.StatusOK, gin.H{"data": info})
}

// POST /token
// Create new token
func CreatePlaidInfo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user := c.PostForm("user")
	itemID := c.PostForm("item_id")
	accessToken := c.PostForm("access_token")

	// Create Book
	token := models.PlaidIntegration{User: user, ItemID: itemID, AccessToken: accessToken, PaymentID: ""}
	db.Create(&token)

	c.JSON(http.StatusOK, gin.H{"data": token})

}

// POST /token/:id
// Find a token
func FindPlaidInfo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	l := c.PostForm("user")
	fmt.Println("l",l)
	// Get model if exist
	var info []models.PlaidIntegration
	if err := db.Select("user", "item_id", "access_token").Where("user = ?", l).First(&info).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info})
}

// DELETE /token/:id
// Delete a token
func DeletePlaidInfo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var info models.PlaidIntegration
	if err := db.Where("id = ?", c.Param("id")).First(&info).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&info)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
