package tjdoe

import (
	"math/rand"

	"github.com/seehuhn/mt19937"
)

type TJDoe struct {
	random *rand.Rand
}

func New(seed int64) *TJDoe {
	tjdoe := new(TJDoe)
	tjdoe.random = rand.New(mt19937.New())
	tjdoe.random.Seed(seed)
	return tjdoe
}

func (tjdoe *TJDoe) AnonymizeDirectory(from, to string, mapping []Mapping) error {
	// TODO
	return nil
}
