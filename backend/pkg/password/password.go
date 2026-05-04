// Package password provides bcrypt hashing and comparison helpers.
package password

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
)

const cost = bcrypt.DefaultCost

// Hash returns the bcrypt hash of the plaintext password.
func Hash(plain string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), cost)
	if err != nil {
		return "", apierror.ErrInternal
	}
	return string(b), nil
}

// Compare checks plain against hash, returning nil on match.
func Compare(hash, plain string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)); err != nil {
		return apierror.New(401, apierror.CodeUnauthorized, "invalid credentials")
	}
	return nil
}
