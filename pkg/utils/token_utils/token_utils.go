package token_utils

import (
	"business-auth/internal/constants"
	"business-auth/internal/dto"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

// Secret key for signing the token
var (
	// Secret key for signing JWT
	secretKey = []byte(viper.GetString("secret.key"))
)

// Claims represents the JWT claims
type Claims struct {
	ClientID string `json:"clientId"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT token
func GenerateJWT(clientId string) (dto.AuthenticationToken, error) {
	expirationTime := time.Now().Add(5 * time.Second)

	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    clientId,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte("your-secret-key")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return dto.AuthenticationToken{}, err
	}

	return dto.AuthenticationToken{Token: signedToken, ExpiredDate: expirationTime.Format(constants.DateTimestampPattern)}, nil
}

// VerifyJWT verifies the JWT token
func VerifyJWT(tokenString string) (*Claims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Check if token is valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
