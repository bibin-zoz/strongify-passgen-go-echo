// models/hash_request.go
package models

// HashRequest holds the input for generating a hash
type HashRequest struct {
	Text       string `json:"text"`
	Secret     string `json:"secret"`
	Length     int    `json:"length"`
	NumSymbols int    `json:"num_symbols"`
	NumNumbers int    `json:"num_numbers"`
}
