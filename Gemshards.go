package gemshards

import (
	"math"
	"time"
)

type generator interface {
	Generate() uint64
	getTime() uint64
}

// Gemshards - unique ID
type Gemshards struct {
	GeneratorID uint64
	Epoch       uint64
	Sequence    uint64
	ID          uint64
}

// Gem - Gemshards generator
// *EppochOffset* - Epoch offset in microseconds (should remain consistent in all instances)
// *GemID* - Must be unique across generator instances; otherwise, IDs will duplicate! Starts at 0
type Gem struct {
	EpochOffset uint64
	time        uint64
	sequence    uint64
	GemID       uint64
}

const totalBits uint64 = 64
const timeBits uint64 = 50                                // Years, should be at least 42 (89 years)
const generatorBits uint64 = 4                            // Generator/Node ID bits
const sequenceBits = totalBits - timeBits - generatorBits // Possible unique IDs per microsecond

var sequenceMax = uint64(math.Pow(2, float64(sequenceBits)))

// Generate - Generates a "shard" - unique id
func (g *Gem) Generate() Gemshards {
	epochTime := g.getTime()
	id := g.GemID<<(totalBits-generatorBits-timeBits) + g.sequence<<(sequenceBits) + epochTime<<(totalBits-timeBits)
	uniqueID := Gemshards{GeneratorID: g.GemID, Epoch: epochTime, ID: id, Sequence: g.sequence}
	return uniqueID
}

func (g *Gem) getTime() uint64 {
	now := uint64(time.Now().UnixNano()/100) - g.EpochOffset

	if g.time == now && g.sequence < sequenceMax {
		g.sequence++
		return g.time
	}
	g.time = now
	g.sequence = 0
	return g.time
}
