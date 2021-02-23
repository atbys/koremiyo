package scraper

import (
	"strconv"
	"strings"

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
		return nil, err
	}
	movie := &domain.Movie{
		Id:       id,
		Title:    getMovieTitle(doc),
		Rate:     getMovieRate(doc),
		Abstruct: "TODO",
		FLink:    targetURL,
		Reviews:  GetMovieReviews(doc),
	}

	return movie, nil
}

func getMovieTitle(doc Document) string {
	movie_titile := doc.Find("div.p-content-detail__main > h2 > span").Text()
	return movie_titile
}

func getMovieRate(doc Document) float64 {
	movie_rate := doc.Find("div.p-content-detail-state > div > div > div.c-rating__score").Text()
	rate, _ := strconv.ParseFloat(movie_rate, 32)
	return rate
}

func GetMovieReviews(doc Document) []string {
	movie_reviews_raw := doc.FindAll("body > div.l-main > div.p-content-detail > div.p-content-detail__foot > div.p-main-area.p-timeline > div.p-mark > div.p-mark__review")
	var movie_reviews []string
	for _, sel := range movie_reviews_raw {
		movie_reviews = append(movie_reviews, sel.Text())
	}
	return movie_reviews
}

func (mrep *MovieRepository) FindClipsByUserId(userId string) ([]int, error) {
	page := 1
	targetURL := baseURL + "/users/" + userId + "/clips" + "?page=" + strconv.Itoa(page)
	doc, err := mrep.GetPage(targetURL)
	numOfClips := getNumOfClips(doc)
	var ids []int

	for numOfClips > 0 {
		clipCountInPage := 0
		res := doc.FindAll("body > div.l-main > div.p-content > div.p-contents-grid > div.c-content-item > a")
		for _, sel := range res {
			l, _ := sel.Attr("href")
			tmp := strings.Split(l, "/")
			id, _ := strconv.Atoi(tmp[len(tmp)-1])
			ids = append(ids, id)
			clipCountInPage += 1
		}
		numOfClips -= clipCountInPage
		page += 1
		targetURL := baseURL + "/users/" + userId + "/clips" + "?page=" + strconv.Itoa(page)
		doc, err = mrep.GetPage(targetURL)
		if err != nil {
			panic(err)
		}
	}

	return ids, err
}

func getNumOfClips(doc Document) int {
	numOfClips_str := doc.Find("body > div.l-main > div.p-users-navi > div > ul > li.p-users-navi__item.p-users-navi__item--clips.is-active > div > span.p-users-navi__count").Text()
	numOfClips, _ := strconv.Atoi(numOfClips_str)

	return numOfClips
}
