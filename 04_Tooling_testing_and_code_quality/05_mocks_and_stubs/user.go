package users

// User represents a simple user record used in examples.
type User struct {
	ID   int
	Name string
}

// UserStore is an interface for retrieving and saving users.
type UserStore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

// GreetUser returns a greeting for a user fetched from the store.
func GreetUser(store UserStore, id int) (string, error) {
	u, err := store.GetUser(id)
	if err != nil {
		return "", err
	}
	return "Hello, " + u.Name, nil
}
