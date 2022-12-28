// Package dao TODO:
package dao

import "gorm.io/gorm"

type ProjectDao struct {
	db *gorm.DB
}

func NewProjectDao(db *gorm.DB) *ProjectDao {
	return &ProjectDao{
		db: db,
	}
}
