package nations

import (
	"durland/internal/domain"
	"durland/internal/registry"
	"math/rand"
)

type ZheleznouhieModifier struct{}

func (z ZheleznouhieModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect
	if ctx.Activity == domain.Zumbalit {
		// Не расходуют удовлетворенность
		result.Satisfaction = 0

		// С вероятностью 0.33 не получают денег от каждой слесандры
		slesandraCount := ctx.Location.GetFaunaCount(domain.Slesandra)
		moneyFromSlesandra := 2.0 * float64(slesandraCount)
		if rand.Float64() < 0.33 {
			result.Money -= moneyFromSlesandra
		}
	}
	return result
}

func init() {
	registry.RegisterNationModifier(domain.Zheleznouhie, ZheleznouhieModifier{})
}
