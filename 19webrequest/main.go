package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	checkError(err)

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // as a caller, always close your connection.

	databytes, err := ioutil.ReadAll(response.Body)

	checkError(err)

	content := string(databytes)

	fmt.Println(content)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
