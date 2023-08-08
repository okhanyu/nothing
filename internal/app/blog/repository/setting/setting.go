package setting

import (
	"errors"
	"log"
	"nothing/internal/app/blog/model/setting"
	repository2 "nothing/internal/app/blog/repository"
	"nothing/internal/pkg/database"
)

type SettingRepository interface {
	FindBatch(sys []int) ([]*setting.SettingBo, error)
	FindByID(int) (*setting.SettingBo, error)
}

type RepositorySettingImpl struct {
	DB *database.DataBase
}

func NewSettingRepository(db *database.DataBase) SettingRepository {
	return &RepositorySettingImpl{DB: db}
}
func (s *RepositorySettingImpl) FindByID(sys int) (*setting.SettingBo, error) {
	var settings *setting.SettingBo
	err := s.DB.GormDB.Table(repository2.SettingsTable).Where("sys = ?", sys).Find(&settings).Error

	if err != nil {
		log.Print(err)
		return nil, err
	}
	return settings, nil
}

func (s *RepositorySettingImpl) FindBatch(sys []int) ([]*setting.SettingBo, error) {
	var settings []*setting.SettingBo

	if sys == nil || len(sys) == 0 {
		return nil, errors.New("参数错误")
	}

	err := s.DB.GormDB.Table(repository2.SettingsTable).Where("id in ?", sys).Find(&settings).Error

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return settings, nil
}
