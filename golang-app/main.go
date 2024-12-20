package main

import (
	"context"
	"fmt"
	"golang-app/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var (
	ctx       = context.Background()
	client, _ = storage.NewClient(ctx, storage.Config{
		Addr:        "redis:6379",
		User:        "default",
		Password:    "my-password",
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 10 * time.Second,
		Timeout:     5 * time.Second,
	})
)

type Capitan struct {
	gorm.Model
	Name      string
	CapNumber string `gorm:"uniqueIndex"`
}

type Ship struct {
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Capitan{})

	db.Create(&Capitan{Name: "Smitt", CapNumber: "A2349GT"})
	db.Create(&Capitan{Name: "Jonson", CapNumber: "B4641KR"})

	return db
}

func main() {
	initDB()

	/*cfg := gorm.Config{
		Addr:        "redis:6379",
		User:        "default",
		Password:    "my-password",
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 10 * time.Second,
		Timeout:     5 * time.Second,
	}

	db, err := gorm.Open(context.Background(), cfg)

	if err != nil {
		panic(err)
	}

	db.Auto*/
	fmt.Println("completed")
}
