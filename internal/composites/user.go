package composites

import (
	"stregy/internal/adapters/api"
	user1 "stregy/internal/adapters/api/user"
	user2 "stregy/internal/adapters/pgorm/user"
	"stregy/internal/domain/user"
)

type UserComposite struct {
	Storage user.Storage
	Service user.Service
	Handler api.Handler
}

func NewUserComposite(composite *PGormComposite) (*UserComposite, error) {
	storage := user2.NewStorage(composite.db)
	service := user.NewService(storage)
	handler := user1.NewHandler(service)
	return &UserComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
