package token_utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Secret key for signing the token
var secretKey = []byte("159357")

// Claims structure to include additional information in the token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateBearerToken(username string) (string, error) {
	// Create a new set of claims
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
		},
	}

	// Create the token with the claims and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
