<p align="center">
<h1 align="center">requestfy</h1>
<p align="center">A library to facilitate the creation of http requests</p>
</p>

## Installation
Use go get command to retrieve the package (Go `^1.19` is required):
```sh
go get -u github.com/RafaelYon/requestfy
```

## Quick Start
 1. Import it in your code:
```go
import "github.com/RafaelYon/requestfy"
```

 2. Set up a HTTP Client:
```go
client := requestfy.NewClient(
    requestfy.ConfigDefault(),
    requestfy.ConfigBaseURL("https://swapi.dev/api/")
)
```

 3. Make a request:
```go
res, err := client.Request().Get("people/1/")
```

## Configuration

### Replacing JSON decoder
When starting the client with the default configuration (`requestfy.ConfigDefault()`) [json.Decoder](https://pkg.go.dev/encoding/json#Decoder) is used.

To replace `json.Decoder` with another implementation it is necessary to specify the do constructor during the creation of the client with the option `ConfigJsonDecoder`.

#### Using go-json
To use "[go-json](https://github.com/goccy/go-json)" just "teach" the client how to build the new decoder:

```go
import "github.com/goccy/go-json"

func main() {
    client := client := requestfy.NewClient(
        requestfy.ConfigDefault(), // Using "ConfigDefault" is optional and its settings may be overwritten by subsequent settings
        requestfy.ConfigJsonDecoder(func(r io.Reader) requestfy.Decoder {
            return json.NewDecoder(r)
        })
    )
}
```