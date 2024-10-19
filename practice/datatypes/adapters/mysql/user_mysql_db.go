package mysql

import (
	"github.com/shivamk2406/Practice/practice/datatypes/interfaces"
	"github.com/shivamk2406/Practice/practice/models"
	"gorm.io/gorm"
)

type mySqlRepository struct {
	db gorm.DB
}

type DB interface {
	interfaces.Repository
}

func NewMysqlRepository(db gorm.DB) DB {
	return &mySqlRepository{
		db: db,
	}
}

func (u *mySqlRepository) GetUserById() models.User {
	return models.User{}
}

func (u *mySqlRepository) CreateUser() error {
	return nil
}

func (u *mySqlRepository) GetUserList(fileName string) []models.User {
	return nil
}
