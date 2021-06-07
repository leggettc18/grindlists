package auth

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/leggettc18/grindlists/api/cache"
	"github.com/twinj/uuid"
)

type TokenDetails struct {
	AccessToken string
	RefreshToken string
	AccessUuid string
	RefreshUuid string
	AtExpires int64
	RtExpires int64
}

func CreateToken(userID int64, secret string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	// Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func CacheAuth(userID int64, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) // Converting Unix to UTC (Time Object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	accessCache, err := cache.NewRedisCacheInstance("access_token", at.Sub(now))
	if err != nil {
		return err
	}
	errAccess := accessCache.Set(td.AccessUuid, strconv.Itoa(int(userID)))
	if errAccess != nil {
		return errAccess
	}

	refreshCache, err := cache.NewRedisCacheInstance("refresh_token", rt.Sub(now))
	if err != nil {
		return err
	}
	errRefresh := refreshCache.Set(td.AccessUuid, strconv.Itoa(int(userID)))
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

type CookieAccess struct {
	Writer http.ResponseWriter
}

// SetToken sets an httponly cookie with a provided token. Intended
// To be used in a resolver which as the CookieAccess struct passed
// into it via the context via a Middleware.
func (access *CookieAccess) SetToken(name string, token string, expiration time.Time) {
	http.SetCookie(access.Writer, &http.Cookie{
		Name: name,
		Value: token,
		HttpOnly: true,
		Path: "/",
		Expires: expiration,
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