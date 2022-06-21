package main

import (
	"log"

	"github.com/shivamk2406/Practice/cmd/app"
)

func main() {
	err := app.Start()
	if err != nil {
		log.Println(err)
	}
}
