package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sqlx.DB
}

type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

func NewDB() (*DB, error) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432  user=postgres dbname=gin_postgresql_sqlx password=12345 sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) GetUsers() ([]User, error) {
	users := []User{}
	err := d.db.Select(&users, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (d *DB) GetUser(id int) (*User, error) {
	user := User{}
	err := d.db.Get(&user, "SELECT id, name, email FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *DB) CreateUser(u *User) error {
	_, err := d.db.NamedExec("INSERT INTO users (name, email) VALUES (:name, :email)", u)
	if err != nil {
		return err
	}
	return nil
}
