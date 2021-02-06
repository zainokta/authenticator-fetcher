package user

//User is user entity
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Phone     string `json:"phone"`
	CreatedAt int64  `json:"created_at"`
}
