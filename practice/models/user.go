package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Country  string `json:"country,omitempty"`
	Jobtitle string `json:"jobtitle,omitempty"`
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetId() string {
	return u.Id
}

func (u User) GetJobTitle() string {
	return u.Jobtitle
}

func (u User) GetCountry() string {
	return u.Country
}

func WriteToJson(data []User, fileName string) error {

	file, _ := json.MarshalIndent(data, "", " ")

	return ioutil.WriteFile(fileName+".json", file, 0644)

}

func GenerateUserData() User {
	err := fake.SetLang("en")
	if err != nil {
		fmt.Println(err.Error())
	}

	return User{
		Id:       uuid.NewString(),
		Name:     fake.FullName(),
		Email:    fake.EmailAddress(),
		Country:  fake.Country(),
		Jobtitle: fake.JobTitle(),
	}
}
