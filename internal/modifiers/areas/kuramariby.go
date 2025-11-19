package areas

import (
	"durland/internal/domain"
	"durland/internal/registry"
	"math/rand"
)

type KuramaribyModifier struct{}

func (k KuramaribyModifier) ModifyEffect(baseEffect domain.Stats, ctx *registry.EffectContext) domain.Stats {
	result := baseEffect
	if ctx.Activity == domain.Gulbonit && ctx.Step > 1 {
		sisandraCount := ctx.Location.GetFaunaCount(domain.Sisandra)
		workingSisandra := 0
		for i := 0; i < sisandraCount; i++ {
			if rand.Float64() >= 0.7 {
				workingSisandra++
			}
		}
		// Пересчитываем удовлетворенность на основе работающих сисяндр
		originalSatisfaction := 2.0 * float64(sisandraCount)
		newSatisfaction := 2.0 * float64(workingSisandra)
		result.Satisfaction += newSatisfaction - originalSatisfaction
	}
	return result
}

func init() {
	registry.RegisterAreaModifier(domain.Kuramariby, KuramaribyModifier{})
}
