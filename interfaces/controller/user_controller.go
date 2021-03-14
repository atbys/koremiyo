package controller

import (
	"github.com/atbys/koremiyo/interfaces/database"
	"github.com/atbys/koremiyo/usecase/user"
)

type UserController struct {
	Interactor user.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: user.UserInteractor{
			Repository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (c *UserController) Create(screenName, filmarksID, password string) *Response {
	res := NewResponse()

	u, err := c.Interactor.Add(screenName, filmarksID, password)
	if err != nil {
		res.Error(err)
		return res
	}

	res.EmbedUserInfo(u)

	return res
}

func (c *UserController) ShowInfo(id int) *Response {
	res := NewResponse()

	u, err := c.Interactor.GetUser(id)
	if err != nil {
		res.Error(err)
		return res
	}

	res.EmbedUserInfo(u)

	return res
}

func (c *UserController) ListFriends(id int) *Response {
	res := NewResponse()

	friends, err := c.Interactor.GetFriends(id)
	if err != nil {
		res.Error(err)
		return res
	}

	res.EmbedUsersInfo(friends)

	return res
}

func (c *UserController) Login(fid, password string, session user.UserSession) *Response {
	res := NewResponse()

	err := c.Interactor.Login(fid, password, session)
	if err != nil {
		res.Error(err)
		return res
	}

	return res
}

func (c *UserController) Logout(session user.UserSession) *Response {
	res := NewResponse()

	err := c.Interactor.Logout(session)
	if err != nil {
		res.Error(err)
		return res
	}

	return res
}

func (c *UserController) SessionCheck(session user.UserSession) *Response {
	res := NewResponse()

	isLoggedin, uid := c.Interactor.SessionCheck(session)
	if isLoggedin {
		res.Message["user_id"] = uid
		res.Message["isLoggedin"] = true
	}

	return res
}
