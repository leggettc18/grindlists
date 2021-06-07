package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userID int64, secret string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

type CookieAccess struct {
	Writer http.ResponseWriter
}

// SetToken sets an httponly cookie with a provided token. Intended
// To be used in a resolver which as the CookieAccess struct passed
// into it via the context via a Middleware.
func (access *CookieAccess) SetToken(token string) {
	http.SetCookie(access.Writer, &http.Cookie{
		Name: "jwtAuth",
		Value: token,
		HttpOnly: true,
		Path: "/",
		Expires: time.Now().Add(time.Hour * 24 * 7),
	})
}

type contextKey struct {
	key string
}

var cookieAccessKeyCtx = contextKey(contextKey{ key: "cookie-access" } )

func setValInCtx (ctx *context.Context, val interface{}) {
	*ctx = context.WithValue(*ctx, cookieAccessKeyCtx, val)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie := CookieAccess{
			Writer: w,
		}
		
		ctx := r.Context()

		// &authCookie is a pointer so anyh changes in the future will change 
		// authCookie
		setValInCtx(&ctx, &authCookie)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetCookieAccess(ctx context.Context) *CookieAccess {
	return ctx.Value(cookieAccessKeyCtx).(*CookieAccess)
}