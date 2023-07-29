package request

import "github.com/golang-jwt/jwt"

type OtpCookieStruct struct {
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
	Email      string `json:"email"`
	Phone      int64  `json:"phone"`
	User_name  string `json:"username"`
	Pass_word  string `json:"password"`
	jwt.StandardClaims
}
