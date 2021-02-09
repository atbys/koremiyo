package presenter

import (
	"github.com/atbys/koremiyo/domain"
	"github.com/atbys/koremiyo/usecase"
)

func (p *HTTPPresenter) ShowUserInfo(u domain.User) (usecase.OutputUserData, error) {
	data := usecase.OutputUserData{
		User: u,
	}
	return data, nil
}

func (p *HTTPPresenter) ShowNull() (usecase.OutputUserData, error) {
	data := usecase.OutputUserData{}
	return data, nil
}
