package main
 
import (
    // "github.com/gin-gonic/gin"
    "medical-test-booking/config"
    "medical-test-booking/controllers"
    "medical-test-booking/server"
    "fmt"
)
 
func main() {
    // Load config
    config.LoadConfig()
 
    dbConfig := server.DatabaseConfig{
        Host:     config.AppConfig.Database.Host,
        Port:     config.AppConfig.Database.Port,
        User:     config.AppConfig.Database.User,
        Password: config.AppConfig.Database.Password,
        Name:     config.AppConfig.Database.Name,
    }
    server.InitDB(dbConfig)
    router := server.InitRouter()
 
    router.POST("/admin/login", controllers.AdminLogin)
    router.POST("/admin/create", controllers.CreateAdmin)
	// router.GET("/api/user/tests", controllers.GetTests)


	//admin routes

	// Get all available tests 
	// Create a new test.
	// Update details of a specific test.
	// Delete a specific test.
	// router.Get("/api/admin/tests", controller.getAllTests) 
	// router.POST("/api/admin/newTest", controller.newTest) 
	// router.PUT("/api/admin/Test/:id", controller.updateTest)
	// router.DELETE("/api/admin/Test/:id", controller.deleteTest)


	//user routes

	//Get all available tests.
	//Get all bookings for the logged-in user.
	//Book a test (test ID, appointment date, etc.).
	//Cancel a specific booking (for the logged-in user)

	router.GET("/api/user/tests", controllers.GetTests) 
	router.GET("/api/user/bookings/:id", controllers.GetBookingByID)
	// router.POST("/api/user/bookTest", controllers.BookTest) 
	// router.PUT("/api/user/booking/:id/cancel", controller.updateTest)

    router.Run(":" + fmt.Sprint(config.AppConfig.Server.Port))
}
 
 


