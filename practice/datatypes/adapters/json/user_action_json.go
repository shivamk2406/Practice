package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	userAction "github.com/shivamk2406/Practice/practice/datatypes/interfaces"
	user "github.com/shivamk2406/Practice/practice/models"
)

type UserActionJson struct{}

func NewUserActionJson() userAction.Repository {

	return &UserActionJson{}

}

func (u *UserActionJson) GetUserById() user.User {
	return user.User{}
}

func (u *UserActionJson) CreateUser() error {
	return nil
}

func (u *UserActionJson) GetUserList(fileName string) []user.User {

	jsonFile, err := os.Open(fileName + ".json")
	if err != nil {
		fmt.Println(err.Error())
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []user.User

	json.Unmarshal(byteValue, &users)

	// for i := 0; i < len(users); i++ {
	// 	fmt.Println("------------------------------------")
	// 	fmt.Println("User Name: " + users[i].GetName())
	// 	fmt.Println("User Id: " + users[i].GetId())
	// 	fmt.Println("User Email: " + users[i].GetEmail())
	// 	fmt.Println("User Job: " + users[i].GetJobTitle())
	// 	fmt.Println("User Country: " + users[i].GetCountry())
	// 	fmt.Println("------------------------------------")
	// }
	return users
}
