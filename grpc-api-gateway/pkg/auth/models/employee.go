package models

type SignUp struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}
