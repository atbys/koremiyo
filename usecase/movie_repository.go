package usecase

import "github.com/atbys/koremiyo/domain"

type MovieRepository interface {
	FindById(int) (*domain.Movie, error)
	FindClipsByUserId(string) ([]int, error)
}

type MutualMovieCache interface {
	Store(List, int) int
	FindById(int) (List, int, error)
}

type MovieOutputPort interface {
	ShowMovieInfo(*domain.Movie) (*OutputData, error)
	ShowIndex(*domain.Movie) (*OutputData, error)
}
