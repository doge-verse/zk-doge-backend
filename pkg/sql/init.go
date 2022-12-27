package sql

import (
	"fmt"
	"github.com/doge-verse/zk-doge-backend/pkg/conf"
	"log"

	"github.com/doge-verse/zk-doge-backend/models"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// Init .
func Init() {
	dbConf := conf.Cfg.Database

	// dsn := dbConf.User + ":" + dbConf.Password + "@tcp(" + dbConf.Host + ":" + cast.ToString(dbConf.Port) + ")/" + dbConf.Dbname + "?" + dbConf.Dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Dbname)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,  // DSN data source name
		DisableDatetimePrecision:  true, // Datetime precision is disabled. Databases before MySQL 5.6do not support it.
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		log.Fatal("init error", err)
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("sqlDB error", err)
		}
		sqlDB.SetMaxIdleConns(dbConf.MaxIdleConn)
		sqlDB.SetMaxOpenConns(dbConf.MaxOpenConn)
		Db = db
	}

	if dbConf.Debug {
		Db = Db.Debug()
	}
	if dbConf.AutoMigrate {
		autoMigrate()
	}
}

func autoMigrate() {
	if err := Db.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Fatal("AutoMigrate error", err)
	}
	log.Print("All table AutoMigrate finish.")
}
