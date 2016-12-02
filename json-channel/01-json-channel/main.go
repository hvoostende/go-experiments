package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type page struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

func (p page) toString() string {
	return toJSON(p)
}

func toJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func getPages() []page {
	raw, err := ioutil.ReadFile("../json/pages.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []page
	json.Unmarshal(raw, &c)
	return c
}

func main() {

	pages := getPages()

	c := make(chan page)

	go func() {
		for _, p := range pages {
			c <- p
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(toJSON(n))
	}

	fmt.Println(toJSON(pages))
}
