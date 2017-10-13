package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func etagFileName() (name string) {
	name = "CURRENT_ETAG.txt"
	return
}

func writefile(fileName string, fileValue []byte) {
	err := ioutil.WriteFile(fileName, fileValue, 0644)
	check(err)
}

func httpEtag(url string) (etag string) {
	resp, err := http.Head(url)

	check(err)
	defer resp.Body.Close()
	etag = resp.Header.Get("etag")
	return
}

func httpBody(url string) (body []byte) {
	resp, err := http.Get("http://example.com/")
	check(err)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	check(err)
	return
}

func readWithPlaceholder(fileName string) (fileValue string) {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		// convert byte array to string
		// https://stackoverflow.com/a/18615786/182484
		contents = []byte("PLACEHOLDER")
	}

	fileValue = string(contents[:])
	return
}

func main() {
	url := "http://www.mtbachelor.com/webcams/snowstake.jpg?44862"
	currentEtag := readWithPlaceholder(etagFileName())
	etagFromSite := httpEtag(url)

	if currentEtag != etagFromSite {
		body := httpBody(url)
		fileName := time.Now().Format("20060102150405") + ".jpg"
		fmt.Println("new file found. saving " + fileName)
		writefile(fileName, body)
		writefile(etagFileName(), []byte(etagFromSite))
	} else {
		fmt.Println("no new file found")
	}
}
