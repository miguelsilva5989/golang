// https://pythonprogramming.net/go/parsing-go-language-programming-tutorial/

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-business-sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	stringBody := string(bytes)
	fmt.Println(stringBody)

	defer resp.Body.Close()

}