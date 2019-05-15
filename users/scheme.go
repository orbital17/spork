package users

type UserID int64

type User struct {
	Id           UserID
	Email        string
	Name         string
	PasswordHash string
}
