package users

import "database/sql"

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (store *Store) AddUser(u User) (id int64, err error) {

	query := `INSERT INTO
	  postgres.public.users (date_added, name, password, email)
	VALUES
	  (
		current_timestamp,
		$1,
		$2,
		$3
	  ) RETURNING id;`
	res := store.db.QueryRow(query, u.Name, u.PasswordHash, u.Email)
	var lastID int64
	err = res.Scan(&lastID)
	return int64(lastID), err
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
		&u.Id,
		&u.Name,
		&u.Email,
		&u.PasswordHash,
	)
	return u, err
}

func (store *Store) GetById(id int64) (User, error) {
	query := `select
		id, name, email, password
	from
		users
	where id = $1; `
	res := store.db.QueryRow(query, id)
	var u User
	err := res.Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.PasswordHash,
	)
	return u, err
}
