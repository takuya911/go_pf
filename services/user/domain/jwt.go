package domain

import (
	"github.com/dgrijalva/jwt-go"
)

// JwtClaims struct
type JwtClaims struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
	jwt.StandardClaims
}

type ctxKey string

// JwtCtxKey const
const JwtCtxKey ctxKey = "jwtClaims"

// TokenType string
type TokenType string

const (
	// IDToken const
	IDToken TokenType = "id_token"
)
