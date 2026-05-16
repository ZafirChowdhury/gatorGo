package main

import (
	"ZafirChowdhury/gatorGo/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error while reading from config file. Error: %s", err)
	}

	cfg.SetUser("zafir")

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error while reading from config file. Error: %s", err)
	}

	fmt.Println(cfg)
}
