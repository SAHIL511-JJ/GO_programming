// Program to read XML file into structure
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Note structure for XML
type Note struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	// sample XML content
	xmlData := `<note>
		<to>Rahul</to>
		<from>Priya</from>
		<heading>Reminder</heading>
		<body>Don't forget the meeting!</body>
	</note>`

	// creating XML file
	file, _ := os.Create("note.xml")
	file.WriteString(xmlData)
	file.Close()

	// reading XML file
	data, err := os.ReadFile("note.xml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// parsing XML into structure
	var note Note
	xml.Unmarshal(data, &note)

	// displaying structure
	fmt.Println("XML Data in Structure:")
	fmt.Println("To:", note.To)
	fmt.Println("From:", note.From)
	fmt.Println("Heading:", note.Heading)
	fmt.Println("Body:", note.Body)
}
