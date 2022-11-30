package shared

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
)

type (
	userClaim struct {
		Name         string `json:"name"`
		encryptedPwd string `json:"encryptedPwd"`
		jwt.RegisteredClaims
	}
)

func GetToken(username, encryptedPwd string) (token string, err error) {
	// Set custom claims, double security

	claims := &userClaim{
		username,
		encryptedPwd,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Hour * 72)},
		},
	}

	// Create token with claims
	// put secret key, i.e = bunga kasih
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("bunga kasih"))
}

func ClaimToken(jwtToken string) (username string, err error) {
	userClaim := userClaim{}
	token, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte("bunga kasih"), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("token is expired")
	}

	return userClaim.Name, nil
}
