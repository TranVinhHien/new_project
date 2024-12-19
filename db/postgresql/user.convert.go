package db

import (
	"context"
	db "new-project/db/sqlc"
	services "new-project/services/entity"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s *SQLStore) GetUser(ctx context.Context, userName string) (services.Users, error) {
	u, err := s.Queries.GetUser(ctx, userName)
	if err != nil {
		return services.Users{}, err
	}
	return u.Convert(), nil
}
func (s *SQLStore) InsertUser(ctx context.Context, user services.Users) (services.Users, error) {
	u, err := s.Queries.CreateUser(ctx, db.CreateUserParams{
		Username: user.Username,
		Password: user.Password,
		FullName: user.FullName,
		CreateAt: user.CreateAt,
	})
	if err != nil {
		return services.Users{}, err
	}
	return u.Convert(), nil
}
func (s *SQLStore) UpdateUser(ctx context.Context, user services.Users) (services.Users, error) {
	u, err := s.Queries.UpdateUser(ctx, db.UpdateUserParams{
		Username: user.Username,
		Password: pgtype.Text{String: user.Password, Valid: user.Password != ""},
		FullName: pgtype.Text{String: user.FullName, Valid: user.FullName != ""},
		IsActive: pgtype.Bool{Bool: user.IsActive},
	})
	if err != nil {
		return services.Users{}, err
	}
	return u.Convert(), nil
}