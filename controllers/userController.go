package controllers
 
import (
    "net/http"
    "github.com/gin-gonic/gin"
    // "golang.org/x/crypto/bcrypt"
    "medical-test-booking/models"
    "medical-test-booking/server"
)
 
func GetTests(c *gin.Context) {
    var tests []models.Test
    // Fetch all tests from the database
    if err := server.DB.Find(&tests).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tests"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"tests": tests})
}

// Struct to represent the booking details with test information
type UserBookingResponse struct {
    BookingID        uint   `json:"booking_id"`
    TestName         string `json:"test_name"`
    Description      string `json:"description"`
    Price            float64 `json:"price"`
    AppointmentDate  string `json:"appointment_date"`
    AppointmentTime  string `json:"appointment_time"`
    Status           string `json:"status"`
}

// getBooking - Retrieves all the bookings for a specific user
func GetBookingByID(c *gin.Context) {
    userID := c.Param("id")  // Get user ID from the URL parameter

    var userBookings []UserBookingResponse

    // Join bookings_tracker and tests table to fetch bookings for a user
    if err := server.DB.Table("bookings_tracker").
        Select("bookings_tracker.id as booking_id, tests.name as test_name, tests.description, tests.price, bookings_tracker.appointment_date, bookings_tracker.appointment_time, bookings_tracker.status").
        Joins("JOIN tests ON bookings_tracker.test_id = tests.id").
        Where("bookings_tracker.user_id = ?", userID).
        Scan(&userBookings).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching bookings"})
        return
    }

    // Check if no bookings found for the user
    if len(userBookings) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "No bookings found for the user"})
        return
    }
    // Return the user bookings in the response
    c.JSON(http.StatusOK, gin.H{"bookings": userBookings})
}

type BookTestInput struct {
    UserID           int       `json:"user_id" binding:"required"`
    TestID           int       `json:"test_id" binding:"required"`
    AppointmentDate  string    `json:"appointment_date" binding:"required"`
    AppointmentTime  string    `json:"appointment_time" binding:"required"`
}

 
 