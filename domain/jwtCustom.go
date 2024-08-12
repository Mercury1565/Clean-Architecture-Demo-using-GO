package domain

import "github.com/dgrijalva/jwt-go"

type JWTCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.StandardClaims
}

type JWTCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
