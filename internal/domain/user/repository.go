package user

type Repository interface {
	GetOne(id string) (*User, error)
	Create(user *User) (*User, error)
	GetByAPIKey(apiKey string) (*User, error)
}
