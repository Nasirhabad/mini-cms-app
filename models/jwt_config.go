package models

type JWTOptions struct {
	SecretKey       string
	ExpiredDuration int
}
