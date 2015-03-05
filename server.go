package main

import (
	//	"encoding/json"
	"fmt"
	"github.com/squeegy06/go-chat-server/objects"
	"net/http"
)

func main() {
	//	http.HandleFunc("/say/", handleSay)
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error")
	}
}

type SayJSON struct {
	Message string
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	user := objects.NewUser()
	err := user.FindByRequest(req)
	if err != nil {
		fmt.Fprintf(w, "No User Found: %s", err.Error())
	} else {
		fmt.Fprintf(w, "Username: %s, UserType: %s, Plain Password: %s, Password: %s, Salt: %s, ID: %d, nameCanonical: %s\n", user.Name, user.GetUserType(), user.PlainPassword, user.Password, user.Salt, user.Id, user.NameCanonical)
		user.Name = "John Snow"
		err = user.Persist()
		if err != nil {
			fmt.Fprintf(w, "Something happened updating your name: %s", err.Error())
		} else {
			fmt.Fprintf(w, "We changed your name")
		}
	}
}

//func handleSay(w http.ResponseWriter, req *http.Request) {
//	if req.Method != "POST" {
//		fmt.Fprintf(w, "You must send a POST")
//		return
//	}
//
//	fmt.Fprintf(w, "You said something\n")
//
//	var message SayJSON
//
//	decoder := json.NewDecoder(req.Body)
//
//	if err := decoder.Decode(&message); err == nil {
//		fmt.Fprintf(w, "Here is what you said: %s\n", message.Message)
//		db, err := sql.Open("mysql", "root:password@/goChatServer")
//		if err != nil {
//			fmt.Println(err.Error())
//			return
//		}
//		defer db.Close()
//
//		insert, err := db.Prepare("INSERT INTO test (message) VALUES(?)")
//		if err != nil {
//			fmt.Println(err.Error())
//			return
//		}
//		defer insert.Close()
//
//		_, err = insert.Exec(message.Message)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//	} else {
//		fmt.Fprintf(w, "An error occured: %s\n", err)
//	}
//}
