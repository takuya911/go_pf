package interactor

import (
	"github.com/takuya911/go_pf/services/user/errors"
	"golang.org/x/crypto/bcrypt"
)

func compareHashAndPass(truePass string, formPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(truePass), []byte(formPass)); err != nil {
		return errors.PasswordFaultError
	}
	return nil
}

func genEncryptedPass(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
