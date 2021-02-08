package usecase

import "github.com/atbys/koremiyo/domain"

type MovieRepository interface {
	FindById(int) (*domain.Movie, error)
}
