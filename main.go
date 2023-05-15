package main

import (
	"log"

	errorpoc "github.com/shivamk2406/Practice/resourcename"
)

func main() {
	//err := graphqlapp.Start()
	err := errorpoc.Start()
	if err != nil {
		log.Println(err)
	}
}
