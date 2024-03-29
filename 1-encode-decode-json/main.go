package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// User to store user information
type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

// UserDb is the database of user struct
type UserDb struct {
	Users []User `json:"users,omitempty"`
	Type  string `json:"type,omitempty"`
}

func main() {
	//createJSONFile()
	f, err := os.Open("user.db.json")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)

	db := UserDb{}
	dec.Decode(&db)
	fmt.Println(db)

}

// createJSONFile is to format the data into json file format
func createJSONFile() {
	users := []User{
		{Username: "Thao Ho", Password: "change me", Email: "thao.ho@gmail.com"},
		{Username: "Thao Nhi", Password: "change me", Email: "thao.ho@gmail.com"},
	}

	db := UserDb{Users: users, Type: "Simple"}

	//fmt.Println(users)
	var buf = new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.Encode(db)
	f, err := os.Create("user.db.json")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(f, buf) // can use the command jq '.' user.db.json in terminal commands to format the json file
}
