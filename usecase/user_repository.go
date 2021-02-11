package usecase

import "github.com/atbys/koremiyo/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindByFid(string) (domain.User, error)
	FindAll() (domain.Users, error)
}
