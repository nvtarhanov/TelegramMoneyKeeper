package model

import (
	"log"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	chatId := 123456

	if err := CreateAccount(chatId); err != nil {
		log.Fatal("Cannot create account")
	}

	// if err := CreateAccount(chatId); err != nil {
	// 	log.Fatal("Failed to create account")
	// }
}
