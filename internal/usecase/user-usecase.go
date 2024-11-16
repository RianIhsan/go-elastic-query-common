package usecase

import (
	"github.com/RianIhsan/go-elastic-query-common/internal/domain"
	"github.com/RianIhsan/go-elastic-query-common/internal/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUseCase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) CreateUser(user domain.User) error {
	return u.repo.Save(user)
}

func (uc *UserUsecase) GetUserByID(id string) (domain.User, error) {
	return uc.repo.FindByID(id)
}
