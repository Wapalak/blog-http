package web

import (
	"github.com/joho/godotenv"
	"log"
)

func (h *Handler) signUp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func (h *Handler) signIn() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
