package main

import (
	"encoding/xml"

	"fmt"
	"log"
	"os"
)


type MarshalerAttr struct {
	XMLName   xml.Name    `xml:"testsuites"`
	Text      string      `xml:",chardata"`
	Name      string      `xml:"name,attr"`
	Testsuite []Testsuite `xml:"testsuite"`
}

type Testsuite struct {
	Text     string     `xml:",chardata"`
	Name     string     `xml:"name,attr"`
	Testcase []Testcase `xml:"testcase"`
}

type Testcase struct {
	Text    string `xml:",chardata"`
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
	File    string `xml:"file,attr"`
	Skipped string `xml:"skipped"`
	Failure string `xml:"failure"`
}


func junit(testsuites, testsuite string) {
	const xmlFile = "./xmlfile.xml"
	// Tcase :=

	// e1 := &MarshalerAttr{Text: "hello", Name: testsuites, Testsuite: {Tsuite.Name: "Testsuite", Testcase:{Tcase.Name: "Testcase"}}

	e1 := &MarshalerAttr{Name: "hello"} 
	e2 := &Testsuite{Name: "Hi"}
	e3 := &Testcase{Name: "case name", File: "test file", Type: "case type"}


	data, err := xml.MarshalIndent(r, " ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("emp.xml", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// f, err := os.Create(xmlFile)
// if nil != err {
// 	log.Fatalln(f)
// }
// defer f.Close()

// xmlEnc := xml.NewEncoder(data)
// xmlEnc.Encode(data)

func main() {
	junit("1", "2")
}
