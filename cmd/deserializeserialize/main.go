package main

import (
	"fmt"
	"github.com/johnfercher/talescoder/pkg/decoder"
	"github.com/johnfercher/talescoder/pkg/encoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACjv369xFJgZGBgYGleWfZa6uaHSdselNJM93fUuQGAIAAKjgjvgoAAAA"

	decoder := decoder.NewDecoder()
	encoder := encoder.NewEncoder()

	slab, err := decoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slabBase64, err := encoder.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
