package users

type UserID int64

type User struct {
	id           UserID
	email        string
	name         string
	passwordHash string
}
