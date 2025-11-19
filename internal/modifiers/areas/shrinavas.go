package main

import (
	"durland/internal/domain"
	"durland/internal/registry"
)

type ShrinavasModifier struct{}

func (s ShrinavasModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect

	if ctx.Activity == domain.Shlyamsat {
		// Добавляет 13% к производительности чучундр
		chuchundraCount := ctx.Location.GetFaunaCount(domain.Chuchundra)
		additionalHealth := 2.0 * float64(chuchundraCount) * 0.13
		result.Health += additionalHealth
	}

	return result
}

func init() {
	registry.RegisterAreaModifier(domain.Shrinavas, ShrinavasModifier{})
}
