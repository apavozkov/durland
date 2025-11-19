package nations

import (
	"durland/internal/domain"
	"durland/internal/registry"
	"math/rand"
)

type MozhoryModifier struct{}

func (m MozhoryModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect
	if ctx.Activity == domain.Gulbonit {
		// Тратят на 23% больше денег при гульбонстве
		result.Money *= 1.23
	}
	if ctx.Activity == domain.Zumbalit && rand.Float64() < 1.0/3.0 {
		// В одном случае из 3 не расходуют здоровье
		result.Health = 0
	}
	return result
}

func init() {
	registry.RegisterNationModifier(domain.Mozhory, MozhoryModifier{})
}
