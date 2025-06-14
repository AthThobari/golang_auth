package utils

import (
  "time"
  "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secretpssword")

/* GenerateToken generates a JWT token with the user ID as 
part of the claims */
func GenerateToken(userID uint) (string, error) {
  claims := jwtMapClaims{}
  claims["user_id"] = userID
  claims["ecp"] = time.Now().Add(time.Hour * 1).Unix()
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(secretKey)
}
