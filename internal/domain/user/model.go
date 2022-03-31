package user

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	PassHash string `json:"owner,omitempty"`
}
