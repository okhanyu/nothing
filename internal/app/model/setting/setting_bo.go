package setting

import (
	"nothing/internal/app/repository"
)

type SettingBo struct {
	ID     int64  `json:"id" form:"id" gorm:"column:id"`
	Type   int    `json:"type" form:"type" gorm:"column:type"`
	Config string `json:"config" form:"config"  gorm:"column:config"`
}

func (SettingBo) TableName() string {
	return repository.SettingsTable
}
