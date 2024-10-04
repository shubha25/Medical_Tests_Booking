package models

import (
    "time"
    // "gorm.io/gorm"
)
 
type Admin struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username" gorm:"unique"`
    Password string `json:"password"`
}


type User struct {
    ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
    Name        string         `json:"name" gorm:"size:255;not null"`
    Email       string         `json:"email" gorm:"size:255;unique;not null"`
    Password    string         `json:"password" gorm:"size:255;not null"`
    PhoneNumber string         `json:"phone_number" gorm:"size:20"`
    Address     string         `json:"address" gorm:"type:text"`
    Role        string         `json:"role" gorm:"size:20;not null"`
    CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

type Test struct {
    ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    Name        string    `json:"name" gorm:"size:255;not null"`
    Description string    `json:"description" gorm:"type:text"`
    Price       float64   `json:"price" gorm:"not null"`
    CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type BookingTracker struct {
    ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    UserID          uint      `json:"user_id" gorm:"not null"`
    TestID          uint      `json:"test_id" gorm:"not null"`
    AppointmentDate time.Time `json:"appointment_date" gorm:"type:date;not null"`
    AppointmentTime time.Time `json:"appointment_time" gorm:"type:time;not null"`
    Status          string    `json:"status" gorm:"size:20;not null"`
    CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}









