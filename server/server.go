package server
 
import (
    "github.com/gin-gonic/gin"
)
 
// InitRouter initializes and returns the Gin router
func InitRouter() *gin.Engine {
    router := gin.Default()
 
    // You can add more middleware here if needed, e.g. CORS, Logging, etc.
 
    return router
}
 