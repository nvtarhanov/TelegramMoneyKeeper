package model

import (
	"log"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	chatId := 12345678

	if err := CreateAccount(chatId); err != nil {
		log.Fatal("Cannot create account")
	}

}

func TestGetAccountById(t *testing.T) {
	chatId := 12345678

	if _, err := GetAccountBySessionID(chatId); err != nil {
		log.Fatal("Cannot find account by id")
	}
}

func TestSetName(t *testing.T) {

	chatId := 12345678
	testName := "test"

	account, err := GetAccountBySessionID(chatId)
	if err != nil {
		log.Fatal("Cannot find account by id")
	}

	if err := SetName(account, testName); err != nil {
		log.Fatal("Cannot set name for account")
	}

}

func TestSetMoneyGoal(t *testing.T) {

	chatId := 12345678
	testMoneyGoal := 99

	account, err := GetAccountBySessionID(chatId)
	if err != nil {
		log.Fatal("Cannot find account by id")
	}

	if err := SetMoneyGoal(account, testMoneyGoal); err != nil {
		log.Fatal("Cannot set money goal for account")
	}

}

func TestSetStartSum(t *testing.T) {

	chatId := 12345678
	testStartSum := 88

	account, err := GetAccountBySessionID(chatId)
	if err != nil {
		log.Fatal("Cannot find account by id")
	}

	if err := SetStartSum(account, testStartSum); err != nil {
		log.Fatal("Cannot set start sum for account")
	}

}
