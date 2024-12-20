package main

import (
	"context"
	"fmt"
	"golang-app/storage"
	"gorm.io/gorm"
	"time"
)

type

func main() {
	cfg := storage.Config{
		Addr:        "redis:6379",
		User:        "default",
		Password:    "my-password",
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 10 * time.Second,
		Timeout:     5 * time.Second,
	}

	db, err := storage.NewClient(context.Background(), cfg)

	if err != nil {
		panic(err)
	}

}
