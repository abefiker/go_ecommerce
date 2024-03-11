package config

import (
    "fmt"
    "github.com/abefiker/go_ecommerce/internals/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var database *gorm.DB
var e error

func Databaseinit() {
    host := "localhost"
    user := "postgres"
    password := "PostgresDb@be1994"
    dbName := "goecommerce"
    port := 5432

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable ", host, user, password, dbName, port)
    database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if e != nil {
        panic("failed to connect database")
    } else {
        fmt.Println("Database is Connected")
        fmt.Println(database)
    }

    if e = database.AutoMigrate(&models.Category{},
        &models.Product{},
        &models.User{},
        &models.Address{},
        &models.Order{},
        &models.OrderItem{},
        &models.CartItem{},
        ); e != nil {
        panic(e)
    }

    // Specify the relationship: a product belongs to a category
    //database.Model(&models.Product{}).BelongsTo(&models.Category{})
}

func DB() *gorm.DB {
    return database
}
