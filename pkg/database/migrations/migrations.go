package database

import (
    "github.com/sixfwa/fiber-gorm/models"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&models.User{})
}

func MigrateUsers(db *gorm.DB) error {
    return db.Migrator().DropTable(&models.User{})
}

func CreateUsersTable(db *gorm.DB) error {
    return db.Migrator().CreateTable(&models.User{})
}