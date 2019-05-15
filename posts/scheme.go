package posts

import "spork/users"

type Post struct {
	Owner users.UserID
	Body  string
}
