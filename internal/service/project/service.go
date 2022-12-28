package project

import (
	"github.com/doge-verse/zk-doge-backend/internal/dao"
	"gorm.io/gorm"
)

type srv struct {
	dao *dao.ProjectDao
}

func newSrv(db *gorm.DB) *srv {
	return &srv{
		dao: dao.NewProjectDao(db),
	}
}
