package scraper

import (
	"strconv"

	"github.com/atbys/koremiyo/domain"
)

type MovieRepository struct {
	Scraper
}

const baseURL = "https://filmarks.com"

func (mrep *MovieRepository) FindById(id int) (*domain.Movie, error) {
	targetURL := baseURL + "/movies/" + strconv.Itoa(id) + "/no_spoiler"
	doc, err := mrep.GetPage(targetURL)

	if err != nil {
		panic(err)
	}
	movie := &domain.Movie{
		Id:       id,
		Title:    getMovieTitle(doc),
		Rate:     0.0,
		Abstruct: "TODO",
		FLink:    "TODO",
		Reviews:  []string{"A", "B"},
	}

	return movie, err
}

func getMovieTitle(doc Document) string {
	movie_titile := doc.Find("div.p-content-detail__main > h2 > span").Text()
	return movie_titile
}
