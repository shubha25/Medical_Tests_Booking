package controllers
 
import (
    "net/http"
    "github.com/gin-gonic/gin"
    // "golang.org/x/crypto/bcrypt"
    "medical-test-booking/models"
    "medical-test-booking/server"
)
 
func AdminLogin(c *gin.Context) {
    var admin models.Admin
    if err := c.ShouldBindJSON(&admin); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
 
    var storedAdmin models.Admin
    if err := server.DB.Where("username = ?", admin.Username).First(&storedAdmin).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username "})
        return
    }
 
    if err := server.DB.Where("password = ?", admin.Password).First(&storedAdmin).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }
 
    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
 
func CreateAdmin(c *gin.Context) {
    var admin models.Admin
    if err := c.ShouldBindJSON(&admin); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := server.DB.Create(&admin).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
        return
    }
    c.JSON(http.StatusCreated, admin)
}
 
 