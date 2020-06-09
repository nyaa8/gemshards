package gemshards

import (
	"testing"
)

func TestGenerators(t *testing.T) {
	gen := Gem{EpochOffset: 15730518251008800, GemID: 0}
	gen2 := Gem{EpochOffset: 15730518251008800, GemID: 10}
	gen3 := Gem{EpochOffset: 15730518251008800, GemID: 15}

	for uniqid := 0; uniqid < 10000000; uniqid++ {
		shard := gen.Generate()
		shard2 := gen2.Generate()
		shard3 := gen3.Generate()

		if shard2.ID == shard.ID || shard2.ID == shard3.ID || shard.ID == shard3.ID {
			t.Errorf("Duplication, iteration number %d, first shard: %d, second shard: %d, third shard: %d", uniqid, shard.ID, shard2.ID, shard3.ID)
		}
	}
}
