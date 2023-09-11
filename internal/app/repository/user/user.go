package user

import (
	"errors"
	user2 "nothing/internal/app/model/user"
	"nothing/internal/app/repository"
	"nothing/internal/pkg/database"
)

type UserRepository interface {
	FindBatch(sys []int) ([]*user2.UserBo, error)
	FindByID(int) (*user2.UserBo, error)
	FindByUsernameAndPassword(username string, password string) (*user2.User, error)
}

type RepositoryUserImpl struct {
	DB *database.DataBase
}

func NewUserRepository(db *database.DataBase) UserRepository {
	return &RepositoryUserImpl{DB: db}
}

func (s *RepositoryUserImpl) FindByUsernameAndPassword(username string, password string) (*user2.User, error) {
	var userObj *user2.User
	err := s.DB.GormDB.Table(repository.UserTable).Where("username = ? and password = ? and deleted = 0", username,
		password).First(&userObj).Error

	if err != nil {
		//log.Print(err)
		return nil, err
	}
	return userObj, nil
}

func (s *RepositoryUserImpl) FindByID(sys int) (*user2.UserBo, error) {
	var users *user2.UserBo
	err := s.DB.GormDB.Table(repository.UserTable).Where("sys = ?", sys).Find(&users).Error

	if err != nil {
		//log.Print(err)
		return nil, err
	}
	return users, nil
}

func (s *RepositoryUserImpl) FindBatch(sys []int) ([]*user2.UserBo, error) {
	var users []*user2.UserBo

	if sys == nil || len(sys) == 0 {
		return nil, errors.New("参数错误")
	}

	err := s.DB.GormDB.Table(repository.UserTable).Where("id in ?", sys).Find(&users).Error

	if err != nil {
		//log.Print(err)
		return nil, err
	}

	return users, nil
}
