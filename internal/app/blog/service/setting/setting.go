package setting

import (
	"log"
	"nothing/internal/app/blog/model/setting"
	settingrepo "nothing/internal/app/blog/repository/setting"
)

type SettingService struct {
	Repository settingrepo.SettingRepository
}

func NewSettingService(repo settingrepo.SettingRepository) *SettingService {
	return &SettingService{
		Repository: repo,
	}
}

func (s *SettingService) FindBatch(sys []int) ([]*setting.SettingBo, error) {
	settings, err := s.Repository.FindBatch(sys)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return settings, nil
}
