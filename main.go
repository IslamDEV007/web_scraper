package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Specify the URL of the web resource we want to download
	url := "https://zerotomastery.io/blog/golang-practice-projects/"

	client := &http.Client{}

	response, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	reader := io.Reader(response.Body)

	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(file, reader)

	file.Close()

}
