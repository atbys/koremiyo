package domain

type User struct {
	ID         int
	Password   string
	ScreenName string
	FilmarksID string
}

type Users []User
