package usecase

import (
	"math/rand"
	"time"

	"github.com/atbys/koremiyo/domain"
)

type MovieInteractor struct {
	MovieRepository MovieRepository
	MovieOutputPort MovieOutputPort
}

type MovieOutputPort interface {
	ShowMovieInfo(*domain.Movie) (*OutputData, error)
	ShowIndex(*domain.Movie) (*OutputData, error)
}

type OutputData struct {
	Config  *OutputConfig
	Movie   *domain.Movie
	Content map[string]string
}

type OutputConfig struct{}

const allMovieNum = 85000

func (interactor *MovieInteractor) GetRecommendation() (*OutputData, error) {
	movie := &domain.Movie{
		Title: "1917",
	}
	return interactor.MovieOutputPort.ShowIndex(movie)
}

func (interactor *MovieInteractor) GetMovieInfo(id int) (*OutputData, error) {
	movie, err := interactor.MovieRepository.FindById(id)
	if err != nil {
		panic(err)
	}
	return interactor.MovieOutputPort.ShowMovieInfo(movie)
}

func randInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func (interactor *MovieInteractor) GetRandom() (*OutputData, error) {
	randId := randInt(allMovieNum)
	movie, _ := interactor.MovieRepository.FindById(randId)
	return interactor.MovieOutputPort.ShowMovieInfo(movie)
}

func (interactor *MovieInteractor) GetRandomFromClips(user_id string) (*OutputData, error) {
	movie := &domain.Movie{}

	return interactor.MovieOutputPort.ShowMovieInfo(movie)
}
