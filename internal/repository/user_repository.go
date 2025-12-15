package repository

import (
	"context"
	"time"

	db "go-projects/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

func (r *UserRepository) CreateUser(ctx context.Context, name string, dob string) (db.User, error) {
	// parse string → time.Time
	parsedDob, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return db.User{}, err
	}

	// convert → pgtype.Date
	pgDob := pgtype.Date{
		Time:  parsedDob,
		Valid: true,
	}

	return r.queries.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  pgDob,
	})
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]db.User, error) {
	return r.queries.ListUsers(ctx)
}
