package presenter

import (
	"github.com/atbys/koremiyo/domain"
	"github.com/atbys/koremiyo/usecase"
)

type HTTPPresenter struct {
	usecase.MovieOutputPort
}

func NewHTTPPresenter() *HTTPPresenter {
	return &HTTPPresenter{}
}

//HTMLで表示するデータ構築をここで行う
func (p *HTTPPresenter) ShowMovieInfo(movie *domain.Movie) (*usecase.OutputData, error) {
	res := &usecase.OutputData{
		Movie:   movie,
		Content: make(map[string]string),
	}
	//データをいい感じに編集したいときはここに書いたりする
	res.Content["movie_title"] = movie.Title
	return res, nil
}

func (p *HTTPPresenter) ShowIndex(movie *domain.Movie) (*usecase.OutputData, error) {
	res := &usecase.OutputData{
		Movie:   movie,
		Content: make(map[string]string),
	}
	res.Content["movie_title"] = movie.Title
	res.Content["page_title"] = "HOGE"

	return res, nil
}
