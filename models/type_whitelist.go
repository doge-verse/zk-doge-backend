package models

type WhiteList struct {
	GormModel
	ProjectID  uint
	Address    string
	Commitment string
}
