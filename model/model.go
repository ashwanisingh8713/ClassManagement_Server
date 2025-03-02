package model

type CreateUser struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
	Password string `json:"password,omitempty"`
}
