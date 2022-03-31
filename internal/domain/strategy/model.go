package strategy

import "stregy/internal/domain/user"

type Strategy struct {
	ID             string
	Name           string
	Description    string
	User           user.User
	Implementation Implementation
}
