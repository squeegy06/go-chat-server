package objects

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	UserType      byte   `json:"-"`
	Id            uint32 `json:"-"`
	Name          string `json:"name,omitempty"`
	NameCanonical string `json:"-"`
	PlainPassword string `json:"plainPassword"`
	Password      string `json:"-"`
	Salt          string `json:"-"`
	LastLogin     string `json:"lastLogin"`
}

//Entity Functions
func (u *User) FindByRequest(req *http.Request) error {
	var decoder *json.Decoder
	var userJSON map[string]*json.RawMessage
	decoder = json.NewDecoder(req.Body)
	_ = decoder.Decode(&userJSON)
	fmt.Println(string(*userJSON["user"]))
	if userJSON["user"] != nil {
		_ = json.Unmarshal(*userJSON["user"], &u) //I don't care if a decoder error occurs.
		fu := u.LoadFromDatabase()
		if fu != nil {
			fmt.Printf(fu.Error())
		}
		return nil
	}
	return errors.New("No user found")
}

func (u *User) GetCacheId() string {
	return ""
}

func (u *User) LoadFromCache() error {
	return nil
}

func (u *User) LoadFromDatabase() error {
	if u.Id > 0 {
		query, err := db.Prepare("SELECT Id, UserType, Name, NameCanonical, Password, Salt, LastLogin FROM users WHERE Id = ?")
		if err != nil {
			panic(err.Error())
		}
		defer query.Close()

		result := query.QueryRow(u.Id)
		_ = result.Scan(&u.Id, &u.UserType, &u.Name, &u.NameCanonical, &u.Password, &u.Salt, &u.LastLogin)
		return nil
	} else {
		u.NameCanonical = Canonicalize(u.Name)
		query, err := db.Prepare("SELECT Id, UserType, Name, NameCanonical, Password, Salt, LastLogin FROM users WHERE nameCanonical = ?")
		if err != nil {
			panic(err.Error())
		}
		defer query.Close()

		result := query.QueryRow(u.NameCanonical)
		fmt.Println(result)
		_ = result.Scan(&u.Id, &u.UserType, &u.Name, &u.NameCanonical, &u.Password, &u.Salt, &u.LastLogin)

		return nil
	}

	return errors.New("No user ID provided")
}

func (u *User) Persist() error {
	//Don't persist anons
	if u.IsType("anon") {
		return errors.New("Cannot persist anons")
	}

	//Update the canonical name.
	u.NameCanonical = Canonicalize(u.Name)

	if len(u.PlainPassword) > 0 {
		//Encrypt the password.
		u.PlainPassword = ""
	}

	if u.Id == 0 {
		query, err := db.Prepare("INSERT INTO users (name , nameCanonical , password , salt , userType) VALUES(?,?,?,?,?)")
		_, err = query.Exec(u.Name, u.NameCanonical, u.Password, u.Salt, u.UserType)
		if err != nil {
			return err
		}
		defer query.Close()
		return nil
	} else {
		query, err := db.Prepare("UPDATE users SET name = ?, nameCanonical = ?, password = ?, salt = ?, userType = ? WHERE id = ?")
		_, err = query.Exec(u.Name, u.NameCanonical, u.Password, u.Salt, u.UserType, u.Id)
		if err != nil {
			return err
		}
		defer query.Close()
		return nil
	}
}

//User Specific Functions
func NewUser() *User {
	return &User{UserType: 2, Name: "anon"}
}

func (u *User) IsType(t string) bool {
	switch u.UserType {
	case 1:
		if strings.EqualFold(t, "user") {
			return true
		}
	case 2:
		if strings.EqualFold(t, "anon") {
			return true
		}
	}
	return false
}

func (u *User) GetUserType() string {
	switch u.UserType {
	case 1:
		return "User"
	case 2:
		return "Anon"
	}
	return "Unknown"
}
