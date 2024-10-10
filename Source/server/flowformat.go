package main

import(

	"encoding/json"
	"encoding/xml"
	"net/http"
//	"strings"
	"fmt"
	"io/ioutil"



)

// Replacement for enums
type FlowType int64

// Replacement for enums

const (
	UnknownFormat FlowType = iota
	JSON
	XML
	HTML
)

func (s FlowType) String() string {
	switch s {
	case JSON:
		return "JSON"
	case XML:
		return "XML"
	case HTML:
		return "HTML"
	default:
		return "Unknown Format"
	}
}
// Replacement for enums


func getURLContent(url string)(string){
	
	
	fmt.Printf("Content of %s ...\n", url)

	resp, err := http.Get(url)

	
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
		
	// do this now so it won't be forgotten
	defer resp.Body.Close()

	// reads html as a slice of bytes
	
	html, err := ioutil.ReadAll(resp.Body)
	
	//if err != nil {
	/*	panic(err)
	}*/
	/*// show the HTML code as a string %s
	fmt.Printf("%s\n", html)
	*/
	return 	fmt.Sprintf("%s", html)

}


func detectFlowFormat(content string) (FlowType, error) {

	//isParsed :=false
	// Generic interface
	var msg interface{}

	// JSON
	if err := json.Unmarshal([]byte(content), &msg);  err == nil{
		return JSON, nil
	}

	// XML
	if err := xml.Unmarshal([]byte(content), &msg); err == nil {
		return XML, nil
	}

	return UnknownFormat,  fmt.Errorf("Unknown Format")
	
}