package main

import (
	"fmt"
	"log"

	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
)

func main() {

	fmt.Println("Test")
	cfg, error := config.Init()

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(cfg)

}
