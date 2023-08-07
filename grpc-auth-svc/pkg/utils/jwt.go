package utils

import (
	"errors"
	"fmt"

	"time"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/domain"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/helper/request"
	"github.com/golang-jwt/jwt"
)

func GenerateTokens(id uint) (string, error) {
	var expiryTime = time.Now().Add(720 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiryTime,
		Id:        fmt.Sprint(id),
	})

	fmt.Println(expiryTime)

	generatedTokens, err := token.SignedString([]byte(config.JwtConfig()))

	return generatedTokens, err

}

func GenerateTokenForOtp(val domain.Employee) (string, error) {
	var expiryTime = time.Now().Add(10 * time.Minute).Unix()
	claims := request.OtpCookieStruct{
		Firstname: val.Firstname,
		Lastname:  val.Lastname,
		Email:     val.Email,
		Phone:     val.Phone,
		Username:  val.Username,
		Password:  val.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.JwtConfig()))

	return token, err

}

func ValidateTokens(signedtoken string) (jwt.StandardClaims, error) {

	token, err := jwt.ParseWithClaims(
		signedtoken, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.JwtConfig()), nil
		})

	if err != nil || !token.Valid {
		return jwt.StandardClaims{}, errors.New("not valid token")
	}

	claim, _ := token.Claims.(*jwt.StandardClaims)

	return *claim, nil
}

func ValidateOtpTokens(signedtoken string) (request.OtpCookieStruct, error) {
	token, err := jwt.ParseWithClaims(
		signedtoken, &request.OtpCookieStruct{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtConfig()), nil
		})

	if err != nil {

		return request.OtpCookieStruct{}, err
	}

	claim, _ := token.Claims.(*request.OtpCookieStruct)

	return *claim, nil
}
