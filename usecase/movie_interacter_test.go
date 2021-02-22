package usecase_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/atbys/koremiyo/domain"
	"github.com/atbys/koremiyo/usecase"
	mock "github.com/atbys/koremiyo/usecase/mock_usecase"
	"github.com/golang/mock/gomock"
)

var mockMovieDB = []domain.Movie{
	{Title: "SE7EN", Reviews: []string{"a", "b"}},
	{Title: "STAND BY ME", Reviews: []string{"c", "d"}},
	{Title: "THE STING", Reviews: []string{"e", "f"}},
	{Title: "SOMMER TIME MACHINE BLUES", Reviews: []string{"g", "h"}},
	{Title: "LOVE LETTER", Reviews: []string{"i", "j"}},
}

var mockUsers = []struct {
	ID    string
	Clips []int
}{
	{ID: "alice", Clips: []int{0, 1, 2, 3, 4}},
	{ID: "bob", Clips: []int{0, 1, 2}},
	{ID: "chris", Clips: []int{2, 3}},
}

func TestGetMovieInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMR := mock.NewMockMovieRepository(mockCtrl)
	mockMOP := mock.NewMockMovieOutputPort(mockCtrl)
	interactor := &usecase.MovieInteractor{
		MovieRepository: mockMR,
		MovieOutputPort: mockMOP,
	}

	mockMR.EXPECT().FindById(1).Return(&mockMovieDB[1], nil)

	want := map[string]interface{}{
		"movie_title":   mockMovieDB[1].Title,
		"movie_reviews": mockMovieDB[1].Reviews,
	}
	mockMOP.EXPECT().ShowMovieInfo(&mockMovieDB[1]).Return(&usecase.OutputData{Msg: want}, nil)

	got, err := interactor.GetMovieInfo(1)

	assertErr(t, err, nil)
	assertMsg(t, got.Msg, want)
}

type MockRNG struct {
	IntnResults []int
}

func (rng *MockRNG) Intn(n int) int {
	result := rng.IntnResults[0]
	rng.IntnResults = rng.IntnResults[1:]
	return result
}

type MSG map[string]interface{}

func TestGetRandomFromClips(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMR := mock.NewMockMovieRepository(mockCtrl)
	mockMOP := mock.NewMockMovieOutputPort(mockCtrl)
	interactor := &usecase.MovieInteractor{
		MovieRepository: mockMR,
		MovieOutputPort: mockMOP,
	}

	testUser := mockUsers[0]

	mockMR.EXPECT().FindClipsByUserId(testUser.ID).Return(testUser.Clips, nil).Times(5) // call 5 times

	usecase.RNG = &MockRNG{IntnResults: []int{0, 1, 2, 3, 4}}
	for i, movie := range mockMovieDB {
		mockMR.EXPECT().FindById(i).Return(&movie, nil)
		want := map[string]interface{}{
			"movie_title":   movie.Title,
			"movie_reviews": movie.Reviews,
		}
		mockMOP.EXPECT().ShowMovieInfo(&movie).Return(&usecase.OutputData{Msg: want}, nil)

		got, err := interactor.GetRandomFromClips(testUser.ID)

		assertErr(t, err, nil)
		assertMsg(t, got.Msg, want)
	}

}

func TestMutualClips(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMR := mock.NewMockMovieRepository(mockCtrl)
	mockMOP := mock.NewMockMovieOutputPort(mockCtrl)
	mockMMC := mock.NewMockMutualMovieCache(mockCtrl)
	interactor := &usecase.MovieInteractor{
		MovieRepository:  mockMR,
		MovieOutputPort:  mockMOP,
		MutualMovieCache: mockMMC,
	}
	var userIDs []string
	for _, user := range mockUsers {
		mockMR.EXPECT().FindClipsByUserId(user.ID).Return(user.Clips, nil)
		userIDs = append(userIDs, user.ID)
	}

	var wantTable []map[string]interface{}
	for i := 0; i < len(mockMovieDB); i++ {
		mockMR.EXPECT().FindById(i).Return(&mockMovieDB[i], nil).AnyTimes()
		wantTable = append(wantTable, map[string]interface{}{"movie_title": mockMovieDB[i].Title, "movie_reviews": mockMovieDB[i].Reviews})
		mockMOP.EXPECT().ShowMovieInfo(&mockMovieDB[i]).Return(&usecase.OutputData{Msg: wantTable[i]}, nil).AnyTimes()
	}

	var cacheList usecase.List
	var cacheIndex int = 0
	var cacheID = -1
	wantRank := []int{2, 3, 1, 0, 4}
	for i := 0; i < len(mockMovieDB); i++ {
		t.Run("get mutual clip "+strconv.Itoa(i), func(t *testing.T) {
			mockMMC.EXPECT().Store(gomock.Any(), gomock.Any()).Do(func(c usecase.List, index int) {
				cacheList = c
				cacheIndex = index
			}).Return(cacheID + 1)
			if cacheID >= 0 {
				mockMMC.EXPECT().FindById(cacheID).Return(cacheList, cacheIndex, nil)
			}

			got, _ := interactor.GetMutualClip(userIDs, cacheID)
			if got.CacheID < 0 {
				t.Error("cache can not execute")
			}
			assertMsg(t, got.Msg, wantTable[wantRank[i]])
			cacheID = got.CacheID
		})
	}
}

func assertMsg(t *testing.T, msg map[string]interface{}, want map[string]interface{}) {
	t.Helper()
	if !reflect.DeepEqual(msg, want) {
		t.Errorf("expected: %v, got: %v", want, msg)
	}
}

func assertErr(t *testing.T, err, want error) {
	t.Helper()
	if err != want {
		t.Errorf("expected: %v, got: %v", want, err)
	}
}
