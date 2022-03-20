package composites

import (
	"stregy/internal/adapters/api"
	user2 "stregy/internal/adapters/api/user"
	user1 "stregy/internal/adapters/pggorm/user"
	"stregy/internal/domain/user"
)

type UserComposite struct {
	Storage user.Storage
	Service user.Service
	Handler api.Handler
}

func NewUserComposite(composite *PGGormComposite) (*UserComposite, error) {
	storage := user1.NewStorage(composite.db)
	service := user.NewService(storage)
	handler := user2.NewHandler(service)
	return &UserComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
