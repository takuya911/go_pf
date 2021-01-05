package interactor

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/takuya911/gqlgen-grpc/services/user/domain"
)

const (
	signKey = "-----BEGIN EC PRIVATE KEY-----\nMIHcAgEBBEIAGT88ebOnAbgmS9Idbns1+VqWV9UN2dvzqiXMmxvAyKNnpoFQxYEL\nLrvmL9uqZaCcbR7EOz5OQyyozKyfqxNiMcigBwYFK4EEACOhgYkDgYYABAB/PCXh\nMMmfHGuR2vm7NLtaa1Jg25CuldjD3LlpFAbrQ0tkfnvskJYRkuFJcbbMGEDLKwvz\nH/HCCi/k8lmynF/DlwH4EAVQTUhkoHO2AUS5zK5oDTKxPN8v86BXBBtbbdVEjZaL\na6hVSC8VOiQD+NeSCWwdN2pY0gYCQHcvxyrCqvAX9Q==\n-----END EC PRIVATE KEY-----"
)

func genTokenPair(id string) (*domain.TokenPair, error) {
	idToken, err := genToken(id, string(domain.IDToken), 3600)
	if err != nil {
		return nil, err
	}
	refreshToken, err := genToken(id, string(domain.RefreshToken), 3600)
	if err != nil {
		return nil, err
	}
	return &domain.TokenPair{
		IDToken:      idToken,
		RefreshToken: refreshToken,
	}, nil
}

func genToken(userID string, sub string, expSec int64) (string, error) {
	expTime := time.Now().Add(time.Duration(expSec) * time.Second)
	claims := &domain.JwtClaims{
		struct {
			ID string `json:"id"`
		}{
			ID: userID,
		},
		jwt.StandardClaims{
			Id:        uuid.New().String(),
			Subject:   sub,
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)

	privateKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(signKey))
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
