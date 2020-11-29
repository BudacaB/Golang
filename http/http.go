package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	str := getContentAsString("https://dog.ceo/api/breeds/image/random")
	fmt.Println("string: ", str)

	dog := getContentAsDogStruct("https://dog.ceo/api/breeds/image/random")
	fmt.Println("struct: ", dog)
	fmt.Println("message: ", dog.Message)
	fmt.Println("status: ", dog.Status)

}

// DogStruct for a dog
type DogStruct struct {
	Message string
	Status  string
}

func getContentAsDogStruct(url string) DogStruct {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	var dog DogStruct

	parsingError := json.Unmarshal(result, &dog)
	if parsingError != nil {
		panic(parsingError)
	}
	return dog
}

func getContentAsString(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
