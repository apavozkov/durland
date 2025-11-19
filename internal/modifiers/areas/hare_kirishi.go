package areas

import (
	"durland/internal/domain"
	"durland/internal/registry"
)

type HareKirishiModifier struct{}

func (h HareKirishiModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect

	// При попадании Дроцентов они расходуют дополнительно по 10% здоровья за каждый интервал
	if durlandets, ok := ctx.Durlandets.(interface {
		GetNation() domain.NationType
		GetStats() domain.Stats
	}); ok {
		if durlandets.GetNation() == domain.Drotsenty {
			healthPenalty := durlandets.GetStats().Health * 0.1
			result.Health -= healthPenalty
		}
	}

	return result
}

func init() {
	registry.RegisterAreaModifier(domain.HareKirishi, HareKirishiModifier{})
}
