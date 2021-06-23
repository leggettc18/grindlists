package auth

import (
	"context"
	"errors"
)

// GetPasswordHash hashes a plaintext password string using argon2 and base64
// encodes it.
func GetPasswordHash(password string) (hashedPassword []byte, err error) {
	params := newDefaultParamsObject()
	hashedPassword, err = generateHash(password, params)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

// VerifyPasswordHash compares a plaintext password and a argon2 hashed and
// base64 encoded password to confirm the plaintext password hashes into the
// same hash.
func VerifyPasswordHash(password string, hashedPassword []byte) (valid bool, err error) {
	return verifyHash(password, string(hashedPassword))
}

func GetUserID(ctx context.Context) (int64, error) {
	user_id, ok := ctx.Value(UserIDKey).(int64)
	if !ok {
		return -1, errors.New("not authenticated")
	}
	return user_id, nil
}