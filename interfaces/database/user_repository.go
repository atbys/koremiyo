package database

import "github.com/atbys/koremiyo/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (first_name, last_name) VALUES (?,?)", u.ScreenName, u.FilmarksID,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, screen_name, filmarks_id FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var screenName string
	var filmarksId string
	row.Next()
	if err = row.Scan(&id, &screenName, &filmarksId); err != nil {
		return
	}
	user.ID = id
	user.ScreenName = screenName
	user.FilmarksID = filmarksId
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, screen_name, filmarks_id FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var screenName string
		var filmarksId string
		if err := rows.Scan(&id, &screenName, &filmarksId); err != nil {
			continue
		}
		user := domain.User{
			ID:         id,
			ScreenName: screenName,
			FilmarksID: filmarksId,
		}
		users = append(users, user)
	}
	return
}