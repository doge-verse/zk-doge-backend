package project

import (
	"github.com/doge-verse/zk-doge-backend/internal/service"
	"gorm.io/gorm"
)

var Srv service.Project

func Init(db *gorm.DB) {
	Srv = newSrv(db)
}
