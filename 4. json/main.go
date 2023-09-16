package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"Coursename"` // it will change the name of the key in json
	Price int
	Platform string
	Password string
	Tags []string
}

func main() {
	fmt.Println("learning json")
	//encodeJson()
	decodeJson()
}

func encodeJson() {
	leoCourses := []course{
		{"react", 20, "youtube", "a1234", []string{"golang", "python"}},
		{"python", 10, "youtube", "b1234", []string{"golang", "python"}},
		{"golang", 30, "youtube", "c1234", []string{"golang", "python"}},
	}

	finalJson, err := json.MarshalIndent(leoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	// whenever a data comes from external source it is in the form of byte.
	jsonData := []byte(`
		{
			"Coursename": "react",
			"Price": 20,
			"Platform": "youtube",
			"Password": "a1234",
			"Tags": [
				"golang",
				"python"
			]
		}
	`)

	var course course
	err := json.Unmarshal(jsonData, &course)

	if err != nil {
		panic(err)
	}

	// %+v will print the key and value of the struct. %v will print only the value.
	fmt.Printf("%+v\n", course)

	// if we want to print the value of a particular key
	var course1 map[string]interface{} // interface{} means any type of value
	err1 := json.Unmarshal(jsonData, &course1)

	if err1 != nil {
		panic(err1)
	}

	for k, v := range course1 {
		fmt.Printf("%s : %v\n", k, v)
	}

}