package main

import (
	"log"

	"github.com/lampjaw/censusgo"
)

func main() {
	b := censusgo.NewQueryBuilder("LampjawScraper", "ps2")
	q := b.NewQuery("character")
	q.Where("name.first_lower").Equals("lampjaw")

	result, _ := q.GetResults()
	log.Printf("%+v", result[0])
}
