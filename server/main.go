package main

import (
	"log"
	"time"

	"github.com/bouhartsev/amonic_airlines/server/internal/app"
)

// @title           Amonic Airlines REST API Documentation
// @version         1.0
// @description     Amonic Airlines fucking project. Go fuck yourself, logach(i).
// @BasePath  /api

func init() {
	loc, err := time.LoadLocation("Europe/Moscow")

	if err != nil {
		log.Fatalf("cannot load location: %v", err)
	}

	time.Local = loc
}

func main() {
	apl, err := app.New()

	if err != nil {
		panic(err)
	}

	if err := apl.Run(); err != nil {
		panic(err)
	}
}
