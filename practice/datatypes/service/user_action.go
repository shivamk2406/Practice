package service

import (
	"errors"
	"strings"

	userAction "github.com/shivamk2406/Practice/practice/datatypes/interfaces"
	user "github.com/shivamk2406/Practice/practice/models"
)

type UserActionService struct {
	repo userAction.UserActionCache
}

func NewUserActionService(repo userAction.UserActionCache) userAction.API {
	return &UserActionService{
		repo: repo,
	}
}

func (u *UserActionService) GetUserByCountry(country string) ([]user.User, error) {
	users := u.repo.GetUserList("test")

	userCountryMap := make(map[string][]user.User)

	for i := 0; i < len(users); i++ {
		_, ok := userCountryMap[strings.ToLower(users[i].GetCountry())]
		if ok {
			userCountryMap[strings.ToLower(users[i].GetCountry())] = append(userCountryMap[strings.ToLower(users[i].GetCountry())], users[i])
		} else {
			usersTemp := make([]user.User, 0)
			usersTemp = append(usersTemp, users[i])
			userCountryMap[strings.ToLower(users[i].GetCountry())] = usersTemp
		}
	}

	users, ok := userCountryMap[strings.ToLower(country)]
	if !ok {
		return nil, errors.New("No User found with given job title")
	}

	return users, nil
}

func (u *UserActionService) GetUserByJobTitle(jobTitle string) ([]user.User, error) {
	users := u.repo.GetUserList("test")

	userJobTitleMap := make(map[string][]user.User)

	for i := 0; i < len(users); i++ {
		_, ok := userJobTitleMap[strings.ToLower(users[i].GetJobTitle())]
		if ok {
			userJobTitleMap[strings.ToLower(users[i].GetJobTitle())] = append(userJobTitleMap[strings.ToLower(users[i].GetJobTitle())], users[i])
		} else {
			usersTemp := make([]user.User, 0)
			usersTemp = append(usersTemp, users[i])
			userJobTitleMap[strings.ToLower(users[i].GetJobTitle())] = usersTemp
		}
	}

	users, ok := userJobTitleMap[strings.ToLower(jobTitle)]
	if !ok {
		return nil, errors.New("No User found with given job title")
	}

	return users, nil
}

func (u *UserActionService) GetUserById(id string) (user.User, error) {
	users := u.repo.GetUserList("test")

	for i := 0; i < len(users); i++ {
		if users[i].GetId() == id {
			return users[i], nil
		}
	}

	return user.User{}, errors.New("No User found with given id")
}
