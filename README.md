# httptracer

Simple HTTP tracer using builtin `httptrace` package.

## Install/Update

```
go get -u github.com/Navid2zp/httptracer
```

### Example

```go
result, err := Tracer("https://google.com", "GET")

fmt.Println("Error: ", err)

fmt.Println("Name Lookup: ", result.NameLookup)
fmt.Println("Connect: ", result.Connect)
fmt.Println("TLS Handshake: ", result.TLSHandshake)
fmt.Println("First Byte: ", result.FirstByte)
fmt.Println("Full Response: ", result.FullResponse)
fmt.Println("Body Size (byte): ", result.BodySize)
```

#### Methods

You can convert results to JSON or XML easier using these methods.

```go
// Returns json encoded bytes
jsonData, _ := result.ToJSON()
fmt.Println(string(jsonData))

// Returns xml encoded bytes
xmlData, _ := result.ToXML()
fmt.Println(string(xmlData))
```
