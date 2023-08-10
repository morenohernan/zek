package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type A struct {
	XMLName xml.Name `xml:"a"`
	Text    string   `xml:",chardata"`
}

// use a.xml for this folder
// cat a.xml | go run xmlreader.go
func main() {
	dec := xml.NewDecoder(os.Stdin)
	var doc A
	if err := dec.Decode(&doc); err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
