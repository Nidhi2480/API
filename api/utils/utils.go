package utils

type User struct {
	Id       any    `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type MobileDetail struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Specs string  `json:"specs"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}
