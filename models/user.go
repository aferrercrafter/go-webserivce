package models

//User basic struct example
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers return all users
func GetUsers() []*User {
	return users
}

//AddUser add user to the list and return the user
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
