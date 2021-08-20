package module

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	userIdSeq int = 1
	users     []*User
)

func addUser(us User) (User, error) {
	u := &User{}
	u.Id = userIdSeq
	u.FirstName = us.FirstName
	u.LastName = us.LastName

	userIdSeq++
	users = append(users, u)
	return *u, nil
}

func getUser(id int) (User, error) {

	if id < userIdSeq {
		return *users[id], nil
	}

	return User{}, nil
}
