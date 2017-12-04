package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	var word string
	var url string

	for _, arg := range os.Args[1:] {
		if isValidURL(arg) {
			url = arg
		} else {
			word = arg
		}
	}

	if url == "" {
		fmt.Println("You must provide at least 1 valid URL!")
		os.Exit(1)
	}

	urlBodyContent := strings.Split(getURLBodyString(url), "\n")

	count := 0

	matchPattern, err := regexp.Compile(word)
	check(err)

	for _, line := range urlBodyContent {
		if matchPattern.MatchString(line) {
			count = count + 1
		}

	}
	fmt.Printf("Occurrance of '%s' on '%s' is %d\n", matchPattern, url, count)

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func isValidURL(url string) bool {
	_, err := http.Get(url)
	return err == nil
}

func getURLBodyString(url string) string {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	return string(body)
}
