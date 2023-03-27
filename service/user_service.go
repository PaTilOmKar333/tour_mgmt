package service

import "tour_mgmt/repo"

type UserServiceInterface interface {
}

type userService struct {
	repo repo.UserRepoInterface
}

func InitUserService(r repo.UserRepoInterface) UserServiceInterface {
	return &userService{
		repo: r,
	}
}
