// models/password_options.go
package models

// PasswordOptions holds the options for generating a password
type PasswordOptions struct {
	Length           int  `json:"length"`
	IncludeNumbers   bool `json:"include_numbers"`
	IncludeSpecials  bool `json:"include_specials"`
	IncludeUppercase bool `json:"include_uppercase"`
	IncludeLowercase bool `json:"include_lowercase"`
}
