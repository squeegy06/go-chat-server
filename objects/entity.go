package objects

import (
	"net/http"
	"strings"
)

type Entity interface {
	FindByRequest(*http.Request) error
	GetCacheId() string
	LoadFromCache() error
	LoadFromDatabase() error
	Persist() error
}

func Canonicalize(str string) string {
	return strings.ToLower(str)
}
