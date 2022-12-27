package user

import (
	"github.com/doge-verse/zk-doge-backend/internal/dao"
	"github.com/doge-verse/zk-doge-backend/models"

	"gorm.io/gorm"
)

type srv struct {
	dao *dao.UserDao
}

func newSrv(db *gorm.DB) *srv {
	return &srv{
		dao: dao.NewUserDao(db),
	}
}

// GetUserByQuery .
func (s *srv) GetUserByQuery(query dao.Query) (*models.User, error) {
	return s.dao.GetUserByQuery(query)
}

// UserRegister .
func (s *srv) UserRegister(user *models.User) (*models.User, error) {
	return s.dao.UserRegister(user)
}

func (s *srv) UpdateEmail(userID uint, email string) error {
	return s.dao.UpdateEmail(userID, email)
}
