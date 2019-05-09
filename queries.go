package spork

func newUser(u User) (id UserID, err error) {

	query := `INSERT INTO
	  postgres.public.users (date_added, name, password, email)
	VALUES
	  (
		current_timestamp,
		$1,
		$2,
		$3
	  ) RETURNING id;`
	db := GetDB()
	res := db.QueryRow(query, u.name, u.passwordHash, u.email)
	var lastID int64
	err = res.Scan(&lastID)
	return UserID(lastID), err
}

func userByEmail(email string) (User, error) {
	query := `select
		id, name, email, password
	from
		users
	where email = $1; `
	db := GetDB()
	res := db.QueryRow(query, email)
	var u User
	err := res.Scan(
		&u.id,
		&u.name,
		&u.email,
		&u.passwordHash,
	)
	return u, err
}
