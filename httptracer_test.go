package httptracer

import (
	"fmt"
	"testing"
)

func TestTracer(t *testing.T) {
	result, err := Tracer("https://google.com", "GET")
	fmt.Println("Error: ", err)
	fmt.Println("Name Lookup: ", result.NameLookup)
	fmt.Println("Connect: ", result.Connect)
	fmt.Println("TLS Handshake: ", result.TLSHandshake)
	fmt.Println("First Byte: ", result.FirstByte)
	fmt.Println("Full Response: ", result.FullResponse)
	fmt.Println("Body Size (byte): ", result.BodySize)

	jsonData, _ := result.ToJSON()
	fmt.Println(string(jsonData))

	xmlData, _ := result.ToXML()
	fmt.Println(string(xmlData))
}