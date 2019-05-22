package users

type User struct {
	Id           int64
	Email        string
	Name         string
	PasswordHash string
}
