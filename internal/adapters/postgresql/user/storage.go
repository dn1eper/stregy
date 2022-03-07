package user

import (
	"context"
	"stregy/internal/domain/user"
	"stregy/pkg/client/postgresql"
)

type storage struct {
	db postgresql.Client
}

func NewStorage(client postgresql.Client) user.Storage {
	return &storage{db: client}
}

func (s *storage) GetOne(ctx context.Context, uuid string) (*user.User, error) {
	panic("not implemented")
}

func (s *storage) GetAll(ctx context.Context, limit, offset int) ([]*user.User, error) {
	// q := `
	// 	SELECT id, name, email FROM users;
	// `

	// rows, err := s.db.Query(ctx, q)
	// if err != nil {
	// 	return nil, err
	// }

	// users := make([]user.User, 0)

	// for rows.Next() {
	// 	var bk User

	// 	err = rows.Scan(&bk.ID, &bk.Name, &bk.Email)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	users = append(users, bk.ToDomain())
	// }

	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }

	// return users, nil
	panic("not implemented")
}

func (s *storage) Create(ctx context.Context, user *user.User) (*user.User, error) {
	panic("not implemented")
}

func (s *storage) Delete(ctx context.Context, user *user.User) error {
	panic("not implemented")
}
