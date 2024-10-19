package memcache

import (
	userAction "github.com/shivamk2406/Practice/practice/datatypes/interfaces"
	user "github.com/shivamk2406/Practice/practice/models"
)

type UserActionCache struct {
	repo userAction.Repository
}

func NewUserActionCache(repo userAction.Repository) userAction.UserActionCache {
	return &UserActionCache{
		repo: repo,
	}
}

func (u *UserActionCache) GetUserById() user.User {
	return u.repo.GetUserById()
}

func (u *UserActionCache) CreateUser() error {
	return u.repo.CreateUser()
}

func (u *UserActionCache) GetUserList(fileName string) []user.User {
	return u.repo.GetUserList(fileName)
}
