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

// func (mrep *MovieRepository) FindByUserId(userId string) ([]string, error) {
// 	page := 1
// 	targetURL := baseURL + "/users/" + userId + "/clips" + "?page=" + strconv.Itoa(page)
// 	doc, err := mrep.GetPage(targetURL)
// 	numOfClips := getNumOfClips(doc)
// 	var ids []string

// 	for numOfClips > 0 {
// 		clipCountInPage := 0
// 		res := doc.Find("body > div.l-main > div.p-content > div.p-contents-grid > div.c-content-item > a")
// 		res.Each(func(i int, sel Selection) {
// 			l, _ := sel.Attr("href")
// 			tmp := strings.Split(l, "/")
// 			id := tmp[len(tmp)-1]
// 			ids = append(ids, id)
// 			clipCountInPage += 1
// 		})
// 		numOfClips -= clipCountInPage
// 		page += 1
// 		targetURL := baseURL + "/users/" + userId + "/clips" + "?page=" + strconv.Itoa(page)
// 		doc, err = mrep.GetPage(targetURL)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	return ids, err
// }

func getNumOfClips(doc Document) int {
	numOfClips_str := doc.Find("body > div.l-main > div.p-users-navi > div > ul > li.p-users-navi__item.p-users-navi__item--clips.is-active > div > span.p-users-navi__count").Text()
	numOfClips, _ := strconv.Atoi(numOfClips_str)

	return numOfClips
}
