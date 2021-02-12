package usecase

import (
	"errors"
	"log"

	"github.com/atbys/koremiyo/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository UserRepository
	UserOutputPort UserOutputPort
}

type UserOutputPort interface {
	ShowUserInfo(u domain.User) (OutputUserData, error)
	ShowNull() (OutputUserData, error)
}

type OutputUserData struct {
	User domain.User
}

func (interactor *UserInteractor) Add(u domain.User) (OutputUserData, error) {
	id, _ := interactor.UserRepository.Store(u)
	u.ID = id
	return interactor.UserOutputPort.ShowUserInfo(u)
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(id int) (OutputUserData, error) {
	user, _ := interactor.UserRepository.FindById(id)
	return interactor.UserOutputPort.ShowUserInfo(user)
}

func (interactor *UserInteractor) GetFriends(id int) ([]domain.User, error) {
	friendIDs, err := interactor.UserRepository.ListFriends(id)
	if err != nil {
		return nil, err
	}
	var friends []domain.User
	for _, userID := range friendIDs {
		user, _ := interactor.UserRepository.FindById(userID)
		friends = append(friends, user)
	}
	return friends, nil
}

func (interactor *UserInteractor) Login(fid string, password string, session UserSession) (*OutputUserData, error) {
	user, _ := interactor.UserRepository.FindByFid(fid)
	savedPassword := user.Password
	inputPassword := password
	// err := compareHashAndPassword(savedPassword, inputPassword)
	// if err != nil {
	// 	return nil, err
	// }
	if savedPassword != inputPassword {
		println(savedPassword)
		println(inputPassword)
		return nil, errors.New("invalid password")
	}
	session.Set("user_id", user.ID)
	session.Save()

	return nil, nil
}

func (interactor *UserInteractor) Logout(session UserSession) error {
	session.Clear()
	session.Save()

	return nil
}

func (interactor *UserInteractor) SessionCheck(session UserSession) (interface{}, error) {
	userID := session.Get("user_id")
	if userID == nil {
		log.Println("not logged in")
		return 0, errors.New("not logged in")
	} else {
		return userID, nil
	}
}

// PasswordEncrypt パスワードをhash化
func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CompareHashAndPassword hashと非hashパスワード比較
func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
