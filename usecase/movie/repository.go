package movie

import "github.com/atbys/koremiyo/domain"

type MovieRepository interface {
	Find(id int) (*domain.Movie, error)
	FindClips(fid string) ([]int, error)
}

type MovieCache interface {
	Store(interface{}) int
	Find(int) (interface{}, error)
}
