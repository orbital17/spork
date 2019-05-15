package users

import "database/sql"

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (store *Store) AddUser(u User) (id UserID, err error) {

	query := `INSERT INTO
	  postgres.public.users (date_added, name, password, email)
	VALUES
	  (
		current_timestamp,
		$1,
		$2,
		$3
	  ) RETURNING id;`
	res := store.db.QueryRow(query, u.name, u.passwordHash, u.email)
	var lastID int64
	err = res.Scan(&lastID)
	return UserID(lastID), err
}

func (store *Store) UserByEmail(email string) (User, error) {
	query := `select
		id, name, email, password
	from
		users
	where email = $1; `
	res := store.db.QueryRow(query, email)
	var u User
	err := res.Scan(
		&u.id,
		&u.name,
		&u.email,
		&u.passwordHash,
	)
	return u, err
}
