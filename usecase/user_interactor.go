package usecase

import (
	"github.com/atbys/koremiyo/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
	UserOutputPort UserOutputPort
}

type UserOutputPort interface {
	ShowUserInfo(u domain.User) (OutputUserData, error)
	ShowNull() (OutputUserData, error)
}

type OutputUserData struct {
	User domain.User
}

func (interactor *UserInteractor) Add(u domain.User) (OutputUserData, error) {
	id, _ := interactor.UserRepository.Store(u)
	u.ID = id
	return interactor.UserOutputPort.ShowUserInfo(u)
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(id int) (OutputUserData, error) {
	user, _ := interactor.UserRepository.FindById(id)
	return interactor.UserOutputPort.ShowUserInfo(user)
}
