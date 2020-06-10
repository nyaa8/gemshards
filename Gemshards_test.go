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

func TestResults(t *testing.T) {
	gen := Gem{EpochOffset: 15730518251008800, GemID: 12}
	arrGen := [100000]uint64{}
	for uniqid := 0; uniqid < 100000; uniqid++ {
		shard := gen.Generate()
		arrGen[uniqid] = shard.ID
		if shard.ID < 0 {
			t.Errorf("Got negative value %d, unsignedwtf", shard.ID)
		}
	}

	for A, a := range arrGen {
		for B, b := range arrGen {
			if a == b && A != B {
				t.Errorf("Duplication found, %d with index %d is equal to %d with index %d", a, A, b, B)
			}
		}
	}
}

