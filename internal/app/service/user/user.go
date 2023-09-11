package user

import (
	"log"
	pkgUser "nothing/internal/app/model/user"
	pkgUserRepo "nothing/internal/app/repository/user"
	"nothing/pkg/jwt"
)

type UserService struct {
	Repository pkgUserRepo.UserRepository
}

func NewUserService(repo pkgUserRepo.UserRepository) *UserService {
	return &UserService{
		Repository: repo,
	}
}

func (s *UserService) FindBatch(sys []int) ([]*pkgUser.UserBo, error) {
	users, err := s.Repository.FindBatch(sys)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return users, nil
}

func (s *UserService) FindByUsernameAndPassword(param pkgUser.UserReq) (*pkgUser.UserBo, error) {
	userPo, err := s.Repository.FindByUsernameAndPassword(param.Username, param.Password)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	token, err := jwt.GenerateToken(param.Username)
	if err != nil {
		return nil, err
	}
	return &pkgUser.UserBo{
		ID:        userPo.ID,
		CreatedAt: userPo.CreatedAt,
		Username:  userPo.Username,
		Role:      userPo.Role,
		Avatar:    userPo.Avatar,
		Token:     token,
	}, nil
}
