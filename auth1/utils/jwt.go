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

/* VerifyToken verifies a token JWT validate */
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
// parse the token
token, err := jwt.Parse(tokenString, 
func(token *jwt.Token) (interface{}, error) {
  // check the signing method
  if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
    return nil, fmt.Errorf("Invalid signing method")
}

return secretKey, nil
})

//check for errors
if err != nil {
  return nil, err
}

//validate the token
if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
  return claims, nil
}

return nil, fmt.Errorf("Invalid token")
}
