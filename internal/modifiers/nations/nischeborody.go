package nations

import (
	"durland/internal/domain"
	"durland/internal/registry"
)

type NischeborodyModifier struct{}

func (n NischeborodyModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect
	if ctx.Activity == domain.Gulbonit {
		// Тратят на 87% меньше денег, но на 76% больше здоровья
		result.Money *= 0.13
		result.Health *= 1.76
	}
	return result
}

func init() {
	registry.RegisterNationModifier(domain.Nischeborody, NischeborodyModifier{})
}
