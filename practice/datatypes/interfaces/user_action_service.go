package interfaces

import (
	user "github.com/shivamk2406/Practice/practice/models"
)

type API interface {
	GetUserByCountry(country string) ([]user.User, error)
	GetUserByJobTitle(jobTitle string) ([]user.User, error)
	GetUserById(id string) (user.User, error)
}
