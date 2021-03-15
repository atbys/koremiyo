package user

import "github.com/atbys/koremiyo/domain"

type UserRepository interface {
	Store(*domain.User) (int, error)
	FindById(int) (*domain.User, error)
	FindByFid(string) (*domain.User, error)
	ListFriends(int) ([]int, error)
	AddFriend(int, int) error
}
