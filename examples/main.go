package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/lampjaw/censusgo"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	b := censusgo.NewQueryBuilder("LampjawScraper", "ps2")
	q := b.NewQuery("character")
	q.Where("name.first_lower").Equals("lampjaw")

	result, _ := q.GetResults()
	log.Printf("%+v", result[0])
}
