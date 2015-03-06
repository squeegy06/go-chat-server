package objects

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

var db *sql.DB

func init() {
	//Load our config file.
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		panic(err.Error())
	}

	var configs map[string]*json.RawMessage
	_ = json.Unmarshal(file, &configs)

	if configs["database"] == nil {
		err = errors.New("Whoops")
		panic(err.Error())
	}

	var dbConfig DatabaseConfig
	_ = json.Unmarshal(*configs["database"], &dbConfig)

	//Build or DSN string
	var dsn []byte
	var value []byte
	if dbConfig.User != "" {
		value = []byte(dbConfig.User)
		dsn = append(dsn, value...)
	}

	if dbConfig.Pass != "" {
		value = []byte(":" + dbConfig.Pass)
		dsn = append(dsn, value...)
	}

	if len(dsn) > 0 {
		value = []byte("@")
		dsn = append(dsn, value...)
	}

	if dbConfig.Host != "" {
		if dbConfig.Port != "" {
			value = []byte("(" + dbConfig.Host + ":" + dbConfig.Port + ")")
			dsn = append(dsn, value...)
		} else {
			value = []byte("(" + dbConfig.Host + ")")
			dsn = append(dsn, value...)
		}
	}

	value = []byte("/")
	dsn = append(dsn, value...)

	if dbConfig.Database != "" {
		value = []byte(dbConfig.Database)
		dsn = append(dsn, value...)
	}

	//Connect to the database
	dsnString := string(dsn)
	db, err = sql.Open("mysql", dsnString)
	if err != nil {
		panic(err.Error())
	}
}
