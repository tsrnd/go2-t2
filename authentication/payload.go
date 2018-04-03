package authentication

import jwt "github.com/dgrijalva/jwt-go"

type payload struct {
	jwt.StandardClaims
	Context map[string]interface{} `json:"context"`
}
