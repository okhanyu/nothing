package setting

import (
	"log"
	"nothing/internal/app/model/setting"
	pkgSettingRepo "nothing/internal/app/repository/setting"
)

type SettingService struct {
	Repository pkgSettingRepo.SettingRepository
}

func NewSettingService(repo pkgSettingRepo.SettingRepository) *SettingService {
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
