package controller

import (
	"net/http"

	"github.com/atbys/koremiyo/domain"
	"github.com/atbys/koremiyo/interfaces/cache"
	"github.com/atbys/koremiyo/interfaces/scraper"
	"github.com/atbys/koremiyo/usecase/movie"
)

type MovieController struct {
	Interactor movie.MovieInteractor
}

func NewMovieController(s scraper.Scraper, c cache.Cache) *MovieController {
	mctrl := &MovieController{
		Interactor: movie.MovieInteractor{
			Repository: &scraper.MovieRepository{
				Scraper: s,
			},
			MovieCache: &cache.MovieCache{
				Cache: c,
			},
		},
	}

	return mctrl
}

type Response struct {
	Status  int
	Message map[string]interface{}
}

func (c *MovieController) Random() *Response {
	res := NewResponse()

	mov, err := c.Interactor.Random()
	if err != nil {
		res.Error(err)
		return res
	}

	res.EmbedMovieInfo(mov)

	return res
}

func (c *MovieController) RandomClip(uid string) *Response {
	res := NewResponse()

	m, err := c.Interactor.RandomFromClips(uid)
	if err != nil {
		res.Error(err)
		return res
	}

	res.EmbedMovieInfo(m)

	return res
}

func (c *MovieController) RandomMutual(uids []string, cacheID int) *Response {
	res := NewResponse()

	m, cacheID, err := c.Interactor.RandomFromMutual(uids, cacheID)
	if err != nil {
		res.Error(err)
		return res
	}

	res.EmbedMovieInfo(m)
	res.Message["cache_id"] = cacheID

	return res
}

func (c *MovieController) MajorityMutual(uids []string, cacheID int) *Response {
	res := NewResponse()

	m, cacheID, err := c.Interactor.MajorityFromMutual(uids, cacheID)
	if err != nil {
		res.Error(err)
		return res
	}

	res.EmbedMovieInfo(m)
	res.Message["cache_id"] = cacheID

	return res
}

func (c *MovieController) Index() *Response {
	res := NewResponse()

	res.Message["recommend"] = "GATTACA"

	return res
}

func NewResponse() *Response {
	res := &Response{
		Status:  http.StatusOK,
		Message: make(map[string]interface{}),
	}

	return res
}

func (res *Response) Error(err error) {
	res.Status = http.StatusBadRequest
	res.Message["error"] = err.Error()
}

func (res *Response) EmbedMovieInfo(m *domain.Movie) {
	res.Message["movie_title"] = m.Title
	res.Message["movie_link"] = m.FLink
	res.Message["movie_rate"] = m.Rate //TODO: N.Nの形式に固定したいので文字列にする
	res.Message["movie_abstruct"] = m.Abstruct
	res.Message["movie_revies"] = m.Reviews
}
