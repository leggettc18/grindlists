package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/leggettc18/grindlists/api/cache"
	"github.com/twinj/uuid"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func createToken(userID int64, secret string) (*TokenDetails, error) {
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

func cacheAuth(userID int64, td *TokenDetails) error {
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
	errRefresh := refreshCache.Set(td.RefreshUuid, strconv.Itoa(int(userID)))
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
		Name:     name,
		Value:    token,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   int(expiration.Unix() - time.Now().Unix()),
	})
}

func (access *CookieAccess) RemoveToken(name string) {
	http.SetCookie(access.Writer, &http.Cookie{
		Name:     name,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	})
}

type contextKey struct {
	key string
}

var cookieAccessKeyCtx = contextKey(contextKey{key: "cookie-access"})
var UserIDKey = contextKey(contextKey{key: "user-id"})
var AccessUuidKey = contextKey(contextKey{key: "access-uuid"})
var RefreshUserIDKey = contextKey(contextKey{"refresh-user-id"})
var RefreshUuidKey = contextKey(contextKey{"refresh-uuid"})

func setValInCtx(ctx *context.Context, val interface{}) {
	*ctx = context.WithValue(*ctx, cookieAccessKeyCtx, val)
}


func getCookieAccess(ctx context.Context) *CookieAccess {
	return ctx.Value(cookieAccessKeyCtx).(*CookieAccess)
}

func extractToken(r *http.Request, name string) (string, error) {
	accessTokenCookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return accessTokenCookie.Value, nil
}

func VerifyToken(r *http.Request, name string, secret string) (*jwt.Token, error) {
	tokenString, err := extractToken(r, name)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request, name string, secret string) error {
	token, err := VerifyToken(r, name, secret)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

type TokenMetaDetails struct {
	Uuid   string
	UserId int64
}

func ExtractTokenMetadata(r *http.Request, tokenType string, secret string) (*TokenMetaDetails, error) {
	token, err := VerifyToken(r, "jwt"+strings.Title(strings.ToLower(tokenType)), secret)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims[strings.ToLower(tokenType)+"_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &TokenMetaDetails{
			Uuid:   accessUuid,
			UserId: userId,
		}, nil
	}
	return nil, err
}

var ErrExpiredToken = errors.New("token expired")

func FetchAuth(ctx context.Context, authD *TokenMetaDetails) (int64, error) {
	accessCache, err := cache.NewRedisCacheInstance("access_token", time.Hour)
	if err != nil {
		return 0, err
	}
	userId, ok := accessCache.Get(ctx, authD.Uuid)
	if !ok {
		return 0, ErrExpiredToken
	}
	userID, err := strconv.ParseInt(userId.(string), 10, 64)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func deleteAuth(prefix string, uuid string) (int64, error) {
	cache, err := cache.NewRedisCacheInstance(prefix, time.Hour)
	if err != nil {
		return 0, err
	}
	deleted, err := cache.Del(uuid)
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
