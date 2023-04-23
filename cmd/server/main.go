package main

import (
	"net/http"

	"github.com/marcioc0sta/finch/configs"
	"github.com/marcioc0sta/finch/internal/entities"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("finch.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.Employee{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", nil)

	http.ListenAndServe(":8000", r)
}
