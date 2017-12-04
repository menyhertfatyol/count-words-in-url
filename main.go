package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	urlBodyContent := strings.Split(getURLBodyString(url), "\n")
	count := 0

	for _, line := range urlBodyContent {
		count = count + countOccurrance(line, word)

	}
	fmt.Printf("Occurrance of '%s' on '%s' is %d\n", word, url, count)

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func countOccurrance(text string, pattern string) int {
	return strings.Count(text, pattern)
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
