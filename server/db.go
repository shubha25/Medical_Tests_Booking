// package server
 
// import (
//     // "fmt"
//     "log"
//     "gorm.io/driver/mysql"
//     "gorm.io/gorm"
//     "medical-test-booking/models"
//     "strconv"
// )
 
// var DB *gorm.DB
 
// // InitDB initializes the database connection
// func InitDB(config struct {
//     Host     string
//     Port     int
//     User     string
//     Password string
//     Name     string
// }) {
//     dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
 
//     var err error
//     DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//     if err != nil {
//         log.Fatalf("Failed to connect to the database: %v", err)
//     }
 
//     // Migrate the schema
//     Migrate()
// }
 
// // Migrate runs the auto migrations for database tables
// func Migrate() {
//     DB.AutoMigrate(&models.Admin{})
//     DB.AutoMigrate(&models.User{})
//     DB.AutoMigrate(&models.Test{})
// }
 
 // server/db.go or server/database.go
package server

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "fmt"
)

// Define DatabaseConfig struct
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Name     string
}

// Initialize the database with the config
var DB *gorm.DB

func InitDB(config DatabaseConfig) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.User, config.Password, config.Host, config.Port, config.Name)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to the database")
    }
    fmt.Println("Database connected!")
}
