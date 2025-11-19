package areas

import (
	"durland/internal/domain"
	"durland/internal/registry"
)

type DolbesburgModifier struct{}

func (d DolbesburgModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect
	if ctx.Activity == domain.Zumbalit {
		// +20% к производительности слесандр
		slesandraCount := ctx.Location.GetFaunaCount(domain.Slesandra)
		result.Money += 2.0 * float64(slesandraCount) * 0.2

		// -30% удовлетворенности (больше тратится)
		result.Satisfaction *= 1.3
	}
	return result
}

func init() {
	registry.RegisterAreaModifier(domain.Dolbesburg, DolbesburgModifier{})
}
