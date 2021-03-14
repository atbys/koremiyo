package controller

import (
	"net/http"
	"strconv"

	"github.com/atbys/koremiyo/domain"
)

type Response struct {
	Status  int
	Message map[string]interface{}
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
	res.Message["movie_rate"] = strconv.FormatFloat(m.Rate, 'f', 1, 64)
	res.Message["movie_abstruct"] = m.Abstruct
	res.Message["movie_revies"] = m.Reviews
}

func (res *Response) EmbedUserInfo(u *domain.User) {
	res.Message["user_screen_name"] = u.ScreenName
	res.Message["user_filmarks_id"] = u.FilmarksID
}

func (res *Response) EmbedUsersInfo(users []*domain.User) {
	screenNames := []string{}
	filmarksIDs := []string{}

	for _, u := range users {
		screenNames = append(screenNames, u.ScreenName)
		filmarksIDs = append(filmarksIDs, u.FilmarksID)
	}

	res.Message["users_screen_name"] = screenNames
	res.Message["users_filmarks_id"] = filmarksIDs
}
