package areas

import (
	"durland/internal/domain"
	"durland/internal/registry"
	"math/rand"
)

type BalbesburgModifier struct{}

func (b BalbesburgModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect
	if ctx.Activity == domain.Zumbalit {
		slesandraCount := ctx.Location.GetFaunaCount(domain.Slesandra)
		for i := 0; i < slesandraCount; i++ {
			if rand.Float64() < 0.15 {
				result.Health -= 0.1
			}
		}
	}
	return result
}

func init() {
	registry.RegisterAreaModifier(domain.Balbesburg, BalbesburgModifier{})
}
