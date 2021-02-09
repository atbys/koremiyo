package presenter

import (
	"strconv"

	"github.com/atbys/koremiyo/domain"
	"github.com/atbys/koremiyo/usecase"
)

//HTMLで表示するデータ構築をここで行う
func (p *HTTPPresenter) ShowMovieInfo(movie *domain.Movie) (*usecase.OutputData, error) {
	res := &usecase.OutputData{
		Movie:   movie,
		Content: make(map[string]string),
	}
	//データをいい感じに編集したいときはここに書いたりする
	res.Content["movie_title"] = movie.Title
	res.Content["movie_rate"] = strconv.FormatFloat(movie.Rate, 'f', 1, 64)
	res.Content["link"] = movie.FLink

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
