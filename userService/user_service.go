package userService

// User does stuff
type User struct{}

// UserService does more stuff
type UserService interface {
	Get(id string) (*User, error)
}
