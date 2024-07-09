// models/hash_request.go
package models

import "github.com/golang-jwt/jwt"

// HashRequest holds the input for generating a hash
type HashRequest struct {
	Text       string `json:"text"`
	PhraseId   int    `json:"phrase_id"`
	Length     int    `json:"length"`
	NumSymbols int    `json:"min_num_symbols"`
	NumNumbers int    `json:"min_num_numbers"`
}
type WordPhrase struct {
	PhraseType string   `json:"phrase_type"`
	Words      []string `json:"words"`
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
type UserDetails struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Phone     string
}
