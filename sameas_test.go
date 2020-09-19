package sameas

import (
	"fmt"
	"testing"

	//"github.com/joker-xii/go-sameas"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	result, err := client.GetResult("http://dbpedia.org/resource/Mogwai")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.URI, "has", result.NumDuplicates, "duplicates:")
	for _, v := range result.Duplicates {
		fmt.Println(v)
	}

}
