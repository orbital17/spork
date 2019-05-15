package users

type UserID int64

type User struct {
	id           UserID `protobuf:"1" json:"id"`
	email        string
	name         string
	passwordHash string
}
