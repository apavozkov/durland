package registry

import "durland/internal/domain"

type EffectContext struct {
	Durlandets interface{}
	Location   *domain.Location
	Activity   domain.ActivityType
	Step       int
}

type EffectModifier interface {
	ModifyEffect(baseEffect domain.Stats, ctx *EffectContext) domain.Stats
}

var (
	nationModifiers = map[domain.NationType]EffectModifier{}
	areaModifiers   = map[domain.AreaType]EffectModifier{}
)

func RegisterNationModifier(nation domain.NationType, modifier EffectModifier) {
	nationModifiers[nation] = modifier
}

func RegisterAreaModifier(area domain.AreaType, modifier EffectModifier) {
	areaModifiers[area] = modifier
}

func ApplyModifiers(baseEffect domain.Stats, ctx *EffectContext, nation domain.NationType, area domain.AreaType) domain.Stats {
	result := baseEffect

	if modifier, exists := nationModifiers[nation]; exists {
		result = modifier.ModifyEffect(result, ctx)
	}
	return result
}

func GetNationModifier(nation domain.NationType) EffectModifier {
	return nationModifiers[nation]
}

func GetAreaModifire(area domain.AreaType) EffectModifier {
	return areaModifiers[area]
}
