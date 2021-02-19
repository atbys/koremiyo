package usecase

import "github.com/atbys/koremiyo/domain"

type MovieRepository interface {
	FindById(int) (*domain.Movie, error)
	FindByUserId(string) ([]int, error)
}

type MutualMovieCache interface {
	Store(List, int) int
	FindById(int) (List, int, error)
}
