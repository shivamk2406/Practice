package primitive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shivamk2406/Practice/practice/datatypes/configs"
	"github.com/shivamk2406/Practice/practice/datatypes/interfaces"

	//practice/datatypes/service/user_action.go
	user "github.com/shivamk2406/Practice/practice/models"
)

func Start(reg interfaces.API) error {
	users := make([]user.User, 0, 100)

	// for i := 0; i < 10000; i++ {
	// 	users = append(users, user.GenerateUserData())
	// }

	// user.WriteToJson(users, "test")

	// for i:=0; i<len(users); i++{
	// 	userInfo:= fmt.Sprintf("Name %s Id %s Email %s",users[i].GetName(),users[i].GetId(),users[i].GetEmail())
	// 	fmt.Println(userInfo)
	// }

	// userAction := userActionJson.NewUserActionJson()
	// userActionCache := userActionCache.NewUserActionCache(userAction)
	// userActionSvc := userActionSvc.NewUserActionService(userActionCache)

	config := configs.LoadAppConfig()
	fmt.Print(config)

	users, err := reg.GetUserByCountry("INDIA")
	if err != nil {
		return err
	}

	fmt.Println(users)
	users, err = reg.GetUserByJobTitle("Sales Associate")
	if err != nil {
		return err
	}

	fmt.Println(users)

	user, err := reg.GetUserById("f46145b1-6111-4b1d-b8ea-b23764735660")
	if err != nil {
		return err
	}

	fmt.Println(user)
	return nil

}

func writeToJson(data []user.User, fileName string) error {
	file, _ := json.MarshalIndent(data, "", " ")

	return ioutil.WriteFile(fileName+".json", file, 0644)

}
