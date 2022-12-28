package models

// Project .
type Project struct {
	GormModel
	Name            string      `json:"name"`
	Address         string      `gorm:"NOT NULL;uniqueIndex:idx_uni;type:varchar(64)" json:"address"`
	Email           string      `json:"email"`
	WhiteList       []WhiteList `json:"whiteList"`
	Condition       string      `json:"condition" gorm:"type:json"`
	AirdropContract string      `json:"airdropContract"`
}
