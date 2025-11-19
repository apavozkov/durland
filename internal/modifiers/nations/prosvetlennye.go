package nations

import (
	"durland/internal/domain"
	"durland/internal/registry"
)

type ProsvetlennyeModifier struct{}

func (p ProsvetlennyeModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect
	if ctx.Activity == domain.Shlyamsat {
		// Дополнительная удовлетворенность от сисяндр в последних 3 локациях
		// Пока просто +1 для примера, можно усложнить логику
		result.Satisfaction += 1.0
	}
	return result
}

func init() {
	registry.RegisterNationModifier(domain.Prosvetlennye, ProsvetlennyeModifier{})
}
