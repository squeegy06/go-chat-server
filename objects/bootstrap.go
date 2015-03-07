package objects

import (
	"database/sql"
	"encoding/json"
	_ "errors"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Pass     string
		Database string
	}
}

var db *sql.DB

func Bootstrap() error {
	//Load our config file.
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		panic(err.Error())
	}

	var Config Config
	_ = json.Unmarshal(file, &Config)

	err = StartDB(Config)

	if err != nil {
		return err
	}

	return nil
}

func StartDB(Config Config) error {
	//Build or DSN string
	var dsn []byte
	var value []byte
	var err error
	if Config.Database.User != "" {
		value = []byte(Config.Database.User)
		dsn = append(dsn, value...)
	}

	if Config.Database.Pass != "" {
		value = []byte(":" + Config.Database.Pass)
		dsn = append(dsn, value...)
	}

	if len(dsn) > 0 {
		value = []byte("@")
		dsn = append(dsn, value...)
	}

	if Config.Database.Host != "" {
		if Config.Database.Port != "" {
			value = []byte("(" + Config.Database.Host + ":" + Config.Database.Port + ")")
			dsn = append(dsn, value...)
		} else {
			value = []byte("(" + Config.Database.Host + ")")
			dsn = append(dsn, value...)
		}
	}

	value = []byte("/")
	dsn = append(dsn, value...)

	if Config.Database.Database != "" {
		value = []byte(Config.Database.Database)
		dsn = append(dsn, value...)
	}

	//Connect to the database
	dsnString := string(dsn)
	db, err = sql.Open("mysql", dsnString)
	if err != nil {
		return err
	}

	return nil
}
