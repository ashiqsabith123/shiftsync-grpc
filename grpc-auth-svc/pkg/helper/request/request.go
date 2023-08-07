package request

import "github.com/golang-jwt/jwt"

type OtpCookieStruct struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	jwt.StandardClaims
}
