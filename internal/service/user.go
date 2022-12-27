package service

import (
	"github.com/doge-verse/zk-doge-backend/internal/dao"
	"github.com/doge-verse/zk-doge-backend/models"
)

type User interface {
	GetUserByQuery(query dao.Query) (*models.User, error)

	UserRegister(user *models.User) (*models.User, error)

	UpdateEmail(userID uint, email string) error
}
