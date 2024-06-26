package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type DB struct {
	db *pgxpool.Pool
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewDB() (*DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "12345", "localhost", "5432", "gin_postgresql_pgxpool")
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %v", err)
	}

	// Optional: Customize pool configuration
	config.MaxConns = 100
	config.MinConns = 1
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	return &DB{db: pool}, nil
}

func (d *DB) GetUsers(ctx context.Context) ([]User, error) {
	users := []User{}

	rows, err := d.db.Query(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (d *DB) GetUser(ctx context.Context, id int) (*User, error) {
	user := User{}
	err := d.db.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *DB) CreateUser(ctx context.Context, u *User) error {
	_, err := d.db.Exec(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", u.Name, u.Email)
	if err != nil {
		return err
	}

	return nil
}
