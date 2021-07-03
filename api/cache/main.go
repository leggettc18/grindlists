package cache

import "github.com/99designs/gqlgen/handler"

type Cache interface {
	handler.PersistedQueryCache
	Set(key string, value interface{}) error
	Del(key string) (int64, error)
}
