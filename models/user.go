package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
