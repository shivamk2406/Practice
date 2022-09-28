package app

import (
	"fmt"
	"log"

	"github.com/shivamk2406/Practice/database"
	"github.com/shivamk2406/Practice/internal/service"

	"github.com/shivamk2406/Practice/configs"
)

func Start() error {
	fmt.Println("App Started")
	conf := configs.LoadAppConfig()

	db, clean, err := database.Open(conf)
	if err != nil {
		log.Println(err)
	}

	user := initializedReg(db)
	reg := service.Registry{UserSvc: user}
	fmt.Println(reg)

	// for i := 0; i < 100; i++ {
	// 	newUser := getUserInput()
	// 	ctx := context.Background()
	// 	userSvc, err := reg.CreateUserSubScription(ctx, newUser)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	log.Printf("Created User is %v", userSvc)
	// }

	defer clean()

	return nil
}

// func getUserInput() *user.Model {
// 	var id string
// 	id = uuid.New().String()
// 	return createUser(id)
// }

// func createUser(id string) *user.Model {
// 	rand.Seed(time.Now().UnixNano())
// 	var constArray = [4]string{
// 		string(constants.Iron),
// 		string(constants.Gold),
// 		string(constants.Silver),
// 		string(constants.Platinum),
// 	}

// 	var namesArray = [30]string{"Shiann", "Frederick", " Rickey ", "Ananda", "Stefani", "Dajuan",
// 		"Jessika", "Nayeli", "Mariana", "Dominick", "Declan", "Tyra", "Raul", "Javonte", "Kimberlee",
// 		"Devon", "Diego", "Joel", "Kameryn", "Mikaela", "Jonah", "Kathy", "Travis", "Darwin", "Camryn",
// 		"Sammy", "Kenan", "Devonte", "Eileen", "Jovanni"}
// 	return &user.Model{ID: id,
// 		Name:         namesArray[rand.Intn(29)+1],
// 		Subscription: constArray[rand.Intn(3)+1]}
// }
