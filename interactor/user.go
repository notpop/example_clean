package interactor

import (
	"example_clean/entity"
	"example_clean/repository"
)

type UserInteractor interface {
	GetUser(id string) (*entity.User, error)
}

type userInteractorImpl struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(ur repository.UserRepository) UserInteractor {
	return &userInteractorImpl{userRepository: ur}
}

func (ui *userInteractorImpl) GetUser(id string) (*entity.User, error) {
	return ui.userRepository.FindByID(id)
}
