package presenter

import (
	"strconv"

	"github.com/atbys/koremiyo/domain"
	"github.com/atbys/koremiyo/usecase"
)

//HTMLで表示するデータ構築をここで行う
func (p *HTTPPresenter) ShowMovieInfo(movie *domain.Movie) (*usecase.OutputData, error) {
	res := &usecase.OutputData{
		Movie: movie,
		Msg:   make(map[string]interface{}),
	}
	//データをいい感じに編集したいときはここに書いたりする
	res.Msg["movie_title"] = movie.Title
	res.Msg["movie_rate"] = strconv.FormatFloat(movie.Rate, 'f', 1, 64)
	res.Msg["filmarks_link"] = movie.FLink
	res.Msg["movie_reviews"] = movie.Reviews

	return res, nil
}

func (p *HTTPPresenter) ShowIndex(movie *domain.Movie) (*usecase.OutputData, error) {
	res := &usecase.OutputData{
		Movie: movie,
		Msg:   make(map[string]interface{}),
	}
	res.Msg["movie_title"] = movie.Title
	res.Msg["page_title"] = "HOGE"

	return res, nil
}
