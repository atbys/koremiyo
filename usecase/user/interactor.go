package user

import (
	"errors"

	"github.com/atbys/koremiyo/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	Repository UserRepository
}

type OutputUserData struct {
	User domain.User
}

func (ui *UserInteractor) Add(screenName, filmarksID, password string) (*domain.User, error) {
	u := &domain.User{
		ScreenName: screenName,
		FilmarksID: filmarksID,
		Password:   password, //TODO: 暗号化
	}

	id, err := ui.Repository.Store(u)
	if err != nil {
		return nil, err
	}

	u.ID = id
	return u, nil
}

func (ui *UserInteractor) GetUser(fid string) (*domain.User, error) {
	u, err := ui.Repository.FindByFid(fid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ui *UserInteractor) GetFriends(id int) ([]*domain.User, error) {
	friendIDs, err := ui.Repository.ListFriends(id)
	if err != nil {
		return nil, err
	}

	var friends []*domain.User
	for _, uid := range friendIDs {
		u, err := ui.Repository.FindById(uid)
		if err != nil {
			return nil, err
		}
		friends = append(friends, u)
	}

	return friends, nil
}

func (ui *UserInteractor) FollowFriend(uid int, friendFID string) error {
	friend, err := ui.GetUser(friendFID) // Followする相手の確認
	if err != nil {
		return err
	}

	friendID := friend.ID
	println("friend id: ", friendID)
	err = ui.Repository.AddFriend(uid, friendID)
	if err != nil {
		return err
	}

	return nil
}

func (ui *UserInteractor) Login(fid, password string, session UserSession) error {
	u, err := ui.Repository.FindByFid(fid)
	if err != nil {
		return err
	}
	savedPassword := u.Password
	inputPassword := password
	//TODO: 暗号化
	if savedPassword != inputPassword {
		return errors.New("invalid password")
	}

	session.Set("user_id", u.ID)
	session.Save()
	return nil
}

func (ui *UserInteractor) Logout(session UserSession) error {
	session.Clear()
	session.Save()

	return nil
}

func (ui *UserInteractor) SessionCheck(session UserSession) (bool, interface{}) {
	//TODO: セッション管理をもっとちゃんとする
	uid := session.Get("user_id")
	if uid == nil {
		return false, nil
	} else {
		return true, uid
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
