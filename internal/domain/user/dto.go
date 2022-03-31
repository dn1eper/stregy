package user

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassHash string `json:"pass_hash"`
}

type UpdateUserDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassHash string `json:"pass_hash"`
}
