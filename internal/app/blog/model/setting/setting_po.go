package setting

import (
	"nothing/internal/app/blog/repository"
)

type SettingPo struct {
	ID     int64  `json:"id" form:"id" gorm:"column:id"`
	Type   int    `json:"type" form:"type" gorm:"column:type"`
	Config string `json:"config" form:"config"  gorm:"column:config"`
}

func (SettingPo) TableName() string {
	return repository.SettingsTable
}
