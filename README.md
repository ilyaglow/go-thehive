go-hive
-------

[TheHive](https://github.com/TheHive-Project/TheHive) client library in Go

It's [early days](https://github.com/ilyaglow/go-hive/issues/1) of the library, so use on your own risk.

## Install

```
go get -u github.com/ilyaglow/go-hive
```

## Usage

```go
package main

import (
	"context"
	"log"

	"github.com/ilyaglow/go-hive"
)

func main() {
	hive, err := thehive.NewClient("https://thehive.domain/", &thehive.ClientOpts{
		Auth: &thehive.APIAuth{
			APIKey: "YOUR-API-KEY",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	cases, _, err := hive.Cases.List(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cases)

	c, _, err := hive.Cases.Get(context.Background(), cases[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(c)
}
```
