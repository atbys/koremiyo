package user_test

import (
	"testing"

	"github.com/atbys/koremiyo/domain"
	"github.com/atbys/koremiyo/usecase/user"
	mock "github.com/atbys/koremiyo/usecase/user/mock"
	"github.com/golang/mock/gomock"
)

var mockUsers = []domain.User{
	{ID: 0, ScreenName: "alice", FilmarksID: "alice_film", Password: "password"},
	{ID: 1, ScreenName: "bob", FilmarksID: "bob_film", Password: "password"},
	{ID: 2, ScreenName: "chris", FilmarksID: "chris_film", Password: "password"},
}

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRep := mock.NewMockUserRepository(ctrl)
	mockRep.EXPECT().Store(gomock.Any()).Return(1, nil)

	ui := user.UserInteractor{mockRep}

	u, err := ui.Add("alice", "alice_film", "password")
	if err != nil {
		t.Fatal("unexpected error")
	}

	if u.ID != 1 || u.Password != "password" {
		t.Errorf("store user is: %v", u)
	}
}

func TestGetFriends(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRep := mock.NewMockUserRepository(ctrl)
	mockRep.EXPECT().ListFriends(0).Return([]int{1, 2}, nil)
	for i := 0; i < len(mockUsers); i++ {
		mockRep.EXPECT().FindById(i).Return(&mockUsers[i], nil).AnyTimes()
	}

	ui := &user.UserInteractor{mockRep}

	friends, err := ui.GetFriends(0)
	if err != nil {
		t.Fatal("unexpected error")
	}

	for id, u := range friends {
		id += 1 // 自分がID:0なので1から友達
		if u.ScreenName != mockUsers[id].ScreenName {
			t.Errorf("got: %v, want: %v", u, mockUsers[id])
		}
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRep := mock.NewMockUserRepository(ctrl)
	mockRep.EXPECT().FindByFid("alice_film").Return(&mockUsers[0], nil).AnyTimes()

	mockSes := NewMockSession()

	ui := &user.UserInteractor{mockRep}

	t.Run("login susess", func(t *testing.T) {
		err := ui.Login("alice_film", "password", mockSes)
		if err != nil {
			t.Error("should be able to login")
		}
	})

	t.Run("login invalid", func(t *testing.T) {
		err := ui.Login("alice_film", "hoge", mockSes)
		if err == nil {
			t.Error("should be not able to login")
		}
	})
}

func TestSessionCheck(t *testing.T) {
	//TODO: やる
}

type MockSession struct {
	Storage map[interface{}]interface{}
}

func NewMockSession() *MockSession {
	return &MockSession{
		Storage: make(map[interface{}]interface{}),
	}
}

func (ms *MockSession) Set(key interface{}, value interface{}) {
	ms.Storage[key] = value
}

func (ms *MockSession) Get(key interface{}) interface{} {
	return ms.Storage[key]
}

func (ms *MockSession) Save() error {
	return nil
}

func (ms *MockSession) Clear() {}
