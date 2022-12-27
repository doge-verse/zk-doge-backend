package dao

import (
	"github.com/doge-verse/zk-doge-backend/models"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

// Query .
type Query struct {
	Limit   uint
	Offset  uint
	UserID  uint
	Email   string
	Address string
}

// Where
func (c Query) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.UserID > 0 {
			db = db.Where("id = ?", c.UserID)
		}
		if len(c.Email) > 0 {
			db = db.Where("email = ?", c.Email)
		}
		if len(c.Address) > 0 {
			db = db.Where("address = ?", c.Address)
		}
		return db
	}
}

// GetUserByQuery .
func (repo *UserDao) GetUserByQuery(query Query) (*models.User, error) {
	user := &models.User{}
	if err := repo.db.Model(user).Scopes(query.where()).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UserRegister .
func (repo *UserDao) UserRegister(user *models.User) (*models.User, error) {
	if err := repo.db.Model(&models.User{}).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserDao) UpdateEmail(userID uint, email string) error {
	if err := repo.db.Model(&models.User{
		GormModel: models.GormModel{
			ID: userID,
		},
	}).Update("email", email).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserDao) count(query Query) (int64, error) {
	var count int64
	if err := repo.db.Model(&models.User{}).Scopes(query.where()).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
