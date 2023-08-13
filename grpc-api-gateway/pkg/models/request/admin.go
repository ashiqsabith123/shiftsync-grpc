package models

type FormApproveID struct {
	FormID int `json:"id"`
}

type FormCorrection struct {
	FormID     int    `json:"empid"`
	Correction string `json:"correction"`
}

type AdminSignUp struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}
