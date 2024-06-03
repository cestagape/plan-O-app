package domain

type User struct {
	ID       int
	Name     string
	Role     Role
	Login    string
	Password string
}
