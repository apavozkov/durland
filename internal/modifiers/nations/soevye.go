package nations

import (
	"durland/internal/domain"
	"durland/internal/registry"
)

type SoyevyeModifier struct{}

func (s SoyevyeModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect
	if ctx.Activity == domain.Zumbalit {
		// Дополнительно 0.12 здоровья на каждую чучундру
		chuchundraCount := ctx.Location.GetFaunaCount(domain.Chuchundra)
		result.Health -= 0.12 * float64(chuchundraCount)
	}
	return result
}

func init() {
	registry.RegisterNationModifier(domain.Soyevye, SoyevyeModifier{})
}
