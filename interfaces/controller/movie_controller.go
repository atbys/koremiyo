package controller

import (
	"net/http"

	"github.com/atbys/koremiyo/interfaces/presenter"
	"github.com/atbys/koremiyo/interfaces/scraper"
	"github.com/atbys/koremiyo/usecase"
)

type MovieController struct {
	Interactor usecase.MovieInteractor
}

func NewMovieController(s scraper.Scraper) *MovieController {
	mctrl := &MovieController{
		Interactor: usecase.MovieInteractor{
			MovieRepository: &scraper.MovieRepository{
				Scraper: s,
			},
			MovieOutputPort: &presenter.HTTPPresenter{},
		},
	}

	return mctrl
}

func (controller *MovieController) Index() (int, *usecase.OutputData) {
	content, _ := controller.Interactor.GetRecommendation()
	return http.StatusOK, content
}

func (controller *MovieController) Random(c Context) {
	movie, err := controller.Interactor.RandomSelect()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, movie)
}
