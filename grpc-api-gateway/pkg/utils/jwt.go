package utils

import (
	"errors"
	"fmt"

	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	"github.com/golang-jwt/jwt"
)

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
