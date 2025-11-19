package nations

import (
	"durland/internal/domain"
	"durland/internal/registry"
)

type DrotsentyModifier struct{}

func (d DrotsentyModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect
	if ctx.Activity == domain.Gulbonit {
		// Вполовину меньше затрат и получаемой удовлетворенности
		result.Health *= 0.5
		result.Money *= 0.5
		result.Satisfaction *= 0.5
	}
	return result
}

func init() {
	registry.RegisterNationModifier(domain.Drotsenty, DrotsentyModifier{})
}
