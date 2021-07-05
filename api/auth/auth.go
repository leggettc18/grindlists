package auth

import (
	"context"
	"errors"
	"net/http"
)

type AuthService interface {
	GetUserID(context.Context) (int64, error)
	AuthMiddleware(http.Handler) (http.Handler)
}

type authSvc struct {
	AccessSecret  string
	RefreshSecret string
}

func NewAuth(aSecret string, rSecret string) AuthService {
	return &authSvc{
		AccessSecret: aSecret,
		RefreshSecret: rSecret,
	}
}

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

func (a *authSvc) GetUserID(ctx context.Context) (int64, error) {
	user_id, ok := ctx.Value(UserIDKey).(int64)
	if !ok {
		return -1, errors.New("not authenticated")
	}
	return user_id, nil
}

func (a *authSvc) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie := CookieAccess{
			Writer: w,
		}

		ctx := r.Context()

		// &authCookie is a pointer so anyh changes in the future will change
		// authCookie
		setValInCtx(&ctx, &authCookie)

		atAuth, err := ExtractTokenMetadata(r, "access", a.AccessSecret)
		if err != nil && err != http.ErrNoCookie {
			// For a no Cookie Error, we want to continue unauthenticated
			// For everything else, we return a bad request status
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		// If tokenAuth came back nil, then the token was not in anyway and
		// none of the below code will run. We skip straight to the resolver
		// or next middleware with no userID in the context.
		if atAuth != nil {
			userID, err := FetchAuth(ctx, atAuth)
			// If any error other than expired token comes up, then there
			// was an error during processing/connecting to redis. We return an
			// Internal Server Error
			if err != nil && err != ErrExpiredToken {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
				// If there was no error, put the userID in the context.
			} else if err == nil {
				ctx = context.WithValue(ctx, UserIDKey, userID)
				ctx = context.WithValue(ctx, AccessUuidKey, atAuth.Uuid)
			}
			// If the error was that the token expired, go on to the
			// next middleware or resolver with no userID in the context.
		}
		rtAuth, err := ExtractTokenMetadata(r, "refresh", a.RefreshSecret)
		if err != nil && err != http.ErrNoCookie {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if rtAuth != nil {
			// We want to let expired tokens through for the resolvers to
			// handle
			if err != nil && err != ErrExpiredToken {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			// Sets separate context key for refresh tokens.
			ctx = context.WithValue(ctx, RefreshUserIDKey, rtAuth.UserId)
			ctx = context.WithValue(ctx, RefreshUuidKey, rtAuth.Uuid)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
