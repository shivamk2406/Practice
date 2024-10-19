package interfaces

import user "github.com/shivamk2406/Practice/practice/models"

type Repository interface {
	GetUserById() user.User
	CreateUser() error
	GetUserList(fileName string) []user.User
}
