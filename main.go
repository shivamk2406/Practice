package main

import (
	"log"

	graphqlapp "github.com/shivamk2406/Practice/graphql/cmd/app"
)

func main() {
	err := graphqlapp.Start()
	if err != nil {
		log.Println(err)
	}
}
