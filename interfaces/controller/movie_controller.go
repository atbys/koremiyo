package controller

import (
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
			Cache: &cache.MovieCache{
				Cache: c,
			},
		},
	}

	return mctrl
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
