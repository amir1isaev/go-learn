package dao

import (
	"go-learn/internal/domain"
)

var UserTableName = AppSchema + "." + "users"

type User struct {
	BaseModel
	Name string `json:"name" db:"name"`
}

func (u User) ToDomain() domain.User {
	return domain.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

var ColumnForCreateUser = []string{
	"name",
}
var ColumnForSelectUser = append(columnForCreateBase,
	"name")

func FromUserDomain(u domain.User) User {
	return User{
		Name: u.Name,
	}
}

func FromUsersDomain(us []domain.User) []User {
	var users []User

	for _, item := range us {
		users = append(users, FromUserDomain(item))
	}

	return users
}
