package app

import (
	"fmt"
	"log"

	"github.com/shivamk2406/Practice/database"

	"github.com/shivamk2406/Practice/configs"
)

func Start() error {
	fmt.Println("App Started")
	conf := configs.LoadAppConfig()

	db, clean, err := database.Open(conf)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(db)
	defer clean()

	return nil
}
