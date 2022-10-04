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