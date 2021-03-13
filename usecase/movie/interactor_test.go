package movie_test

import (
	"reflect"
	"testing"

	"github.com/atbys/koremiyo/domain"
	"github.com/atbys/koremiyo/usecase/movie"
	"github.com/golang/mock/gomock"

	mock "github.com/atbys/koremiyo/usecase/movie/mock"
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

type MockRNG struct {
	Stock []int
}

func (m *MockRNG) Intn(max int) int {
	ret := m.Stock[0]
	m.Stock = m.Stock[1:]
	return ret
}

func TestRandom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRep := mock.NewMockMovieRepository(ctrl)
	mockRep.EXPECT().Find(0).Return(&mockMovieDB[0], nil)
	mockRep.EXPECT().Find(1).Return(&mockMovieDB[1], nil)
	mockRep.EXPECT().Find(2).Return(&mockMovieDB[2], nil)

	movie.RNG = &MockRNG{Stock: []int{0, 1, 2}}
	mi := movie.MovieInteractor{mockRep, nil}

	for i := 0; i < 3; i++ {
		m, err := mi.Random()
		if err != nil {
			t.Error("unexpected error")
		}

		if m.Title != mockMovieDB[i].Title {
			t.Errorf("want: %v, got: %v", mockMovieDB[i].Title, m.Title)
		}
	}
}

func TestRandomFromClips(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRep := mock.NewMockMovieRepository(ctrl)
	mockRep.EXPECT().FindClips("alice").Return(mockUsers[0].Clips, nil).AnyTimes()
	for i := 0; i < 5; i++ {
		mockRep.EXPECT().Find(i).Return(&mockMovieDB[i], nil)
	}

	movie.RNG = &MockRNG{Stock: []int{0, 1, 2, 3, 4}}
	mi := movie.MovieInteractor{mockRep, nil}

	for i := 0; i < 5; i++ {
		m, err := mi.RandomFromClips("alice")
		if err != nil {
			t.Fatal("unexpected error")
		}

		if m.Title != mockMovieDB[i].Title {
			t.Errorf("want: %v, got: %v", mockMovieDB[i].Title, m.Title)
		}
	}
}

func TestRandomFromMutual(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRep := mock.NewMockMovieRepository(ctrl)
	mockCache := mock.NewMockMovieCache(ctrl)
	mi := movie.MovieInteractor{mockRep, mockCache}
	mockRep.EXPECT().FindClips("chris").Return([]int{2, 3}, nil)
	mockRep.EXPECT().FindClips("bob").Return([]int{0, 1, 2}, nil)
	for i := 0; i < len(mockMovieDB); i++ {
		mockRep.EXPECT().Find(i).Return(&mockMovieDB[i], nil).AnyTimes()
	}
	mockCache.EXPECT().Store(gomock.Any()).Return(1)
	//mockCache.EXPECT().Find(1).Return(gomock.Any()).AnyTimes()
	movie.RNG = &MockRNG{Stock: []int{0, 1, 2, 3, 4}}
	t.Run("no cache", func(t *testing.T) {
		m, _, err := mi.RandomFromMutual([]string{"chris", "bob"}, -1)
		if err != nil {
			t.Fatal("unexpected error")
		}
		mid := m.ID

		if !inIntSlice(mid, []int{0, 1, 2, 3}) {
			t.Errorf("unexpected value is got: %v", m)
		}
	})
	//cacheありのテスト書くのめんどくさい
}

func TestMajorityFromMutual(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRep := mock.NewMockMovieRepository(ctrl)
	mockCache := mock.NewMockMovieCache(ctrl)
	mi := movie.MovieInteractor{mockRep, mockCache}
	mockRep.EXPECT().FindClips("chris").Return([]int{2, 3}, nil)
	mockRep.EXPECT().FindClips("bob").Return([]int{0, 1, 2}, nil)
	for i := 0; i < len(mockMovieDB); i++ {
		mockRep.EXPECT().Find(i).Return(&mockMovieDB[i], nil).AnyTimes()
	}
	mockCache.EXPECT().Store(gomock.Any()).Return(1)

	t.Run("no cache", func(t *testing.T) {
		m, _, err := mi.MajorityFromMutual([]string{"chris", "bob"}, -1)
		if err != nil {
			t.Fatal("unexpected error")
		}

		if !reflect.DeepEqual(m, &mockMovieDB[2]) {
			t.Errorf("want: %v, got: %v", mockMovieDB[2].Title, m.Title)
		}
	})

	//cacheありのテスト
}

func inIntSlice(value int, s []int) bool {
	for _, i := range s {
		if i == value {
			return true
		}
	}

	return false
}

func assertErr(t *testing.T, err, want error) {
	t.Helper()
	if err != want {
		t.Errorf("expected: %v, got: %v", want, err)
	}
}
