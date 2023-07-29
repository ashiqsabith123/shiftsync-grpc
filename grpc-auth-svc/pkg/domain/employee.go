package domain

type Employee struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
