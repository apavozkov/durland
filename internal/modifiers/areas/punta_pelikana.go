package areas

import (
	"durland/internal/domain"
	"durland/internal/registry"
	"math/rand"
)

type PuntaPelikanaModifier struct{}

func (p PuntaPelikanaModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect

	// Эффект применяется начиная со 2 интервала нахождения в локации
	if ctx.Step >= 2 {
		if ctx.Activity == domain.Gulbonit {
			// Сисяндры генерируют на 23% больше удовлетворенности
			sisandraCount := ctx.Location.GetFaunaCount(domain.Sisandra)
			additionalSatisfaction := 2.0 * float64(sisandraCount) * 0.23
			result.Satisfaction += additionalSatisfaction

			// С вероятностью 0.2 списывается 50% всех денег
			if rand.Float64() < 0.2 {
				if durlandets, ok := ctx.Durlandets.(interface {
					GetStats() domain.Stats
				}); ok {
					result.Money -= durlandets.GetStats().Money * 0.5
				}
			}
		}
	}

	return result
}

func init() {
	registry.RegisterAreaModifier(domain.PuntaPelikana, PuntaPelikanaModifier{})
}
