// models/hash_request.go
package models

import "github.com/golang-jwt/jwt"

// HashRequest holds the input for generating a hash
type HashRequest struct {
	Text       string `json:"text"`
	Secret     string `json:"secret"`
	Length     int    `json:"length"`
	NumSymbols int    `json:"num_symbols"`
	NumNumbers int    `json:"num_numbers"`
}
type Claims struct {
	ID       uint
	Username string
	Email    string
	jwt.StandardClaims
}
type TokenUser struct {
	Users        Claims
	AccessToken  string
	RefreshToken string
}
