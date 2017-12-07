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
	words := []string{}
	urls := []string{}

	for _, arg := range os.Args[1:] {
		if isValidURL(arg) {
			urls = append(urls, arg)
		} else {
			words = append(words, arg)
		}
	}

	if len(urls) == 0 {
		fmt.Println("You must provide at least 1 valid URL!")
		os.Exit(1)
	}

	for _, url := range urls {
		urlBodyContent := strings.Split(getURLBodyString(url), "\n")
		count := 0

		for _, word := range words {
			matchPattern, err := regexp.Compile(word)
			check(err)

			for _, line := range urlBodyContent {
				if matchPattern.MatchString(line) {
					count++
				}
			}
			fmt.Printf("'%s' has %d occurrance(s) of '%s'\n", url, count, matchPattern)
		}
	}

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
