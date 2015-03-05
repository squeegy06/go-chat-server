package objects

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

var db *sql.DB

func init() {
	//Start our database connection for this package.
	var err error
	db, err = sql.Open("mysql", "root:password@/goChatServer")
	if err != nil {
		panic(err.Error())
	}
}

func Canonicalize(str string) string {
	return strings.ToLower(str)
}
