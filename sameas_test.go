package sameas

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	//"github.com/joker-xii/go-sameas"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	result, err := client.GetResult("http://dbpedia.org/resource/Mogwai")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, int64(50), result.NumDuplicates)
	assert.Equal(t, "http://www.wikidata.org/entity/Q645980", result.URI)

	fmt.Println(result.URI, "has", result.NumDuplicates, "duplicates:")
	for _, v := range result.Duplicates {
		fmt.Println(v)
	}

}
