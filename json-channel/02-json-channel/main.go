package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//create struct

type page struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	//read json file
	bs := getJSON("../json/pages.json")

	//create channel	5t64rt6y12 ---------------------------
	c := make(chan byte)

	//put json on channel
	go func() {
		for _, v := range bs {
			c <- v
		}
		close(c)
	}()

	var receivedJSON []byte

	//receive from channel
	for n := range c {
		receivedJSON = append(receivedJSON, n)
	}

	//json to struct
	var page []page

	json.Unmarshal(receivedJSON, &page)
	fmt.Println(page)

}

func getJSON(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return content
}
