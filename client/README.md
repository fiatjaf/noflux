Noflux API Client
===================

[![PkgGoDev](https://pkg.go.dev/badge/github.com/fiatjaf/noflux/client)](https://pkg.go.dev/github.com/fiatjaf/noflux/client)

Client library for Noflux REST API.

Installation
------------

```bash
go get -u github.com/fiatjaf/noflux/client
```

Example
-------

```go
package main

import (
	"fmt"
	"os"

	noflux "github.com/fiatjaf/noflux/client"
)

func main() {
    // Authentication with username/password:
    client := noflux.New("https://api.example.org", "admin", "secret")

    // Authentication with an API Key:
    client := noflux.New("https://api.example.org", "my-secret-token")

    // Fetch all feeds.
    feeds, err := client.Feeds()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(feeds)

    // Backup your feeds to an OPML file.
    opml, err := client.Export()
    if err != nil {
        fmt.Println(err)
        return
    }

    err = os.WriteFile("opml.xml", opml, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
}
```
