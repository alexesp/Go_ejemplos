package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Structura para reprsentacion total slice
type Users struct {
	Users []User `json:"users"`
}

// Representacion interna user
type User struct {
	Name string `json: "name"`
	Type string `json: "type"`
	Job  string `json: "job"`
}
type Professor struct {
	Name       string     `json:"name"`
	ScienceID  int        `json:"science_id"`
	IsWorking  bool       `json:"is_working"`
	University University `json:"university"`
}
type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Job: %s\n", u.Job)
}
func main() {
	prof1 := Professor{
		Name:      "Bob",
		ScienceID: 81263123,
		IsWorking: true,
		University: University{
			Name: "BMSTU",
			City: "Moscow",
		},
	}
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}

	var users Users

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(byteValue)
	json.Unmarshal(byteValue, &users)

	byteArr, err := json.MarshalIndent(prof1, "@", "####")
	for _, u := range users.Users {
		PrintUser(&u)
	}

	//file, err := os.Create("test.txt")

	//file, err = os.Open("test.txt")

	//err = file.Truncate(0)
	//file.Write([]byte("Un texto de prueba"))
	err = os.WriteFile("test.txt", byteArr, 0666)
	//err = os.WriteFile("test.txt", byteArr, 0664)
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()
	//fmt.Println(users)

}
