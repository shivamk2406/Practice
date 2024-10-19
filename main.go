package main

import (
	"log"

	datatypes "github.com/shivamk2406/Practice/practice/datatypes"
)

func main() {
	err := datatypes.Start()
	if err != nil {
		log.Println(err)
	}
}
