package user

import (
	"log"
	"nothing/internal/app/blog/model/user"
	userrepo "nothing/internal/app/blog/repository/user"
	"nothing/pkg/jwt"
)

type UserService struct {
	Repository userrepo.UserRepository
}

func NewUserService(repo userrepo.UserRepository) *UserService {
	return &UserService{
		Repository: repo,
	}
}

func (s *UserService) FindBatch(sys []int) ([]*user.UserBo, error) {
	users, err := s.Repository.FindBatch(sys)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return users, nil
}

func (s *UserService) FindByUsernameAndPassword(param user.UserReq) (*user.UserBo, error) {
	userPo, err := s.Repository.FindByUsernameAndPassword(param.Username, param.Password)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	token, err := jwt.GenerateToken(param.Username)
	if err != nil {
		return nil, err
	}
	return &user.UserBo{
		ID:        userPo.ID,
		CreatedAt: userPo.CreatedAt,
		Username:  userPo.Username,
		Role:      userPo.Role,
		Avatar:    userPo.Avatar,
		Token:     token,
	}, nil
}
