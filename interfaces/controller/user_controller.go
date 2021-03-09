package controller

import (
	"net/http"

	"github.com/atbys/koremiyo/domain"
	"github.com/atbys/koremiyo/interfaces/database"
	"github.com/atbys/koremiyo/interfaces/presenter"
	"github.com/atbys/koremiyo/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
			UserOutputPort: &presenter.HTTPPresenter{},
		},
	}
}

func (controller *UserController) Create(u domain.User) (int, usecase.OutputUserData) {
	content, err := controller.Interactor.Add(u)
	if err != nil {
		return http.StatusBadGateway, usecase.OutputUserData{}
	}
	return http.StatusOK, content
}

func (controller *UserController) Index() {
	_, err := controller.Interactor.Users()
	if err != nil {
		return
	}
}

func (controller *UserController) Show(id int) (int, usecase.OutputUserData) {
	content, err := controller.Interactor.UserById(id)
	if err != nil {
		return http.StatusBadGateway, usecase.OutputUserData{}
	}
	return http.StatusOK, content
}

func (controller *UserController) SelectFriend(id int) (int, []domain.User) {
	users, err := controller.Interactor.GetFriends(id)
	if err != nil {
		return http.StatusBadGateway, nil
	}
	return http.StatusOK, users
}

func (controller *UserController) SignUp(u domain.User) int {
	_, err := controller.Interactor.Add(u)
	if err != nil {
		return http.StatusBadGateway
	}
	return http.StatusOK
}

func (controller *UserController) Login(fid string, password string, session usecase.UserSession) int {
	_, err := controller.Interactor.Login(fid, password, session)
	if err != nil {
		return http.StatusBadGateway
	}
	return http.StatusOK
}

func (controller *UserController) Logout(session usecase.UserSession) int {
	err := controller.Interactor.Logout(session)
	if err != nil {
		return http.StatusBadGateway
	}
	return http.StatusOK
}

func (controller *UserController) SessionCheck(session usecase.UserSession) (bool, interface{}) {
	userID, err := controller.Interactor.SessionCheck(session)
	if err != nil {
		return false, nil
	}
	return true, userID
}
