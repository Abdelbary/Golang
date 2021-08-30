package module

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	userIdSeq int = 1
	users     []*User
)

func AddUser(us User) (User, error) {
	if us.Id != 0 {
		return User{}, errors.New("New users must not include Ids")
	}
	u := &User{}
	u.Id = userIdSeq
	userIdSeq++
	u.FirstName = us.FirstName
	u.LastName = us.LastName

	users = append(users, u)
	return *u, nil
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if id == u.Id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with ID '%v' not found ", id)

}

func GetUserByID(id int) (User, error) {

	if id < userIdSeq {
		return *users[id], nil
	}

	return User{}, nil
}

func GetUsers() (users []*User) {
	return users
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.Id == u.Id {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.Id)
}
