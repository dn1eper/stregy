package user

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassHash string `json:"pass_hash"`
}

type UpdateUserDTO struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassHash int    `json:"pass_hash"`
}
