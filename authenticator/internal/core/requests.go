package core

type UserRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
