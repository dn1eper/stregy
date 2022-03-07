package user

type User struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	PassHash string `json:"owner,omitempty"`
}
