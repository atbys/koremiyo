package presenter

import "github.com/atbys/koremiyo/usecase"

type HTTPPresenter struct {
	usecase.MovieOutputPort
}

func NewHTTPPresenter() *HTTPPresenter {
	return &HTTPPresenter{}
}
