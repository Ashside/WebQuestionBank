package api

const (
	STUDENT = iota
	TEACHER
	ADMIN
)

type Account struct {
	Username string
	Password string
	Role     int
}

func (a *Account) Login() bool {
	return true
}
