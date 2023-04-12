package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("This is another lesson on JSON. A bit more depth!!!")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	spheresEDcourses := []course{
		{"ReactJS Bootcamp", 299, "SpheresED.dev", "spheresed123", []string{"web-dev", "code"}},
		{"MERN Bootcamp", 199, "SpheresED.dev", "mern123", []string{"mern", "development"}},
		{"React Native Bootcamp", 499, "SpheresED.dev", "rn123", []string{"mobile", "cross platform"}},
		{"GoLang Bootcamp", 99, "SpheresED.dev", "123go", nil},
	}

	//package data as JSON data

	finalJson, err := json.MarshalIndent(spheresEDcourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "SpheresED.dev",      
		"tags": ["web-dev","code"]
	}
	`)

	var spheresEDcourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &spheresEDcourse)
		fmt.Printf("%#v\n", spheresEDcourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// Sometimes, we may just want to add data to key value. In such cases, we take the approach...

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n ", k, v, v)
	}
}
