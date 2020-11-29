package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dataContainer := getData()
	for _, dc := range dataContainer.Data {
		dc.Content = getContent(dc.Link)
	}

	output, mErr := json.Marshal(dataContainer)
	handleError(mErr)
	writeFile(output)
}

func getData() DataContainer {
	res, err := ioutil.ReadFile("../data.json")
	handleError(err)

	var dataContainer DataContainer
	jErr := json.Unmarshal(res, &dataContainer)
	handleError(jErr)

	return dataContainer
}

func getContent(path string) string {
	res, err := ioutil.ReadFile(path)
	handleError(err)
	return string(res)
}

func writeFile(content []byte) {
	writePerm := os.FileMode(0755)
	wErr := ioutil.WriteFile("../dataChanged.json", content, writePerm)
	handleError(wErr)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
