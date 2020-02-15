package main

import (
	"log"

	"github.com/lampjaw/censusgo"
)

func main() {
	b := censusgo.NewQueryBuilder("example", "ps2")
	q := b.NewQuery("character")
	q.Where("name.first_lower").Equals("lampjaw")

	result, err := q.GetResults()
	if err != nil {
		log.Printf("%v", err)
		return
	}

	log.Printf("%s", result)
}
