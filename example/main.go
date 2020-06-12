package main

import (
	"fmt"

	"nyaa.science/gemshards"
)

func main() {
	generator := gemshards.Gem{EpochOffset: 5730518316208800, GemID: 0}
	shard := generator.Generate()
	fmt.Println("Unique ID:", shard.ID)
}
