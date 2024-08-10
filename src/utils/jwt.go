package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mojtabafarzaneh/social_media/src/types"
)

var SecretKey = []byte("most_valuable_secret_of_mine")

func GenerateToken(ttl time.Duration, user types.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["isAdmin"] = user.IsAdmin

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(token string, signedJWTKey string) (jwt.MapClaims, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims, nil
}
