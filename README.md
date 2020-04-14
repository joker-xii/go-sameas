# go-sameas

A golang client for [sameas.org](http://sameas.org/)

## Example

```go
package main

import (
	"fmt"
	"github.com/joker-xii/go-sameas"
)

func main() {
	client := sameas.NewClient()
	result, err := client.GetResult("http://dbpedia.org/resource/Mogwai")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.URI, "has", result.NumDuplicates, "duplicates:")
	for _, v := range result.Duplicates {
		fmt.Println(v)
	}

}

```

