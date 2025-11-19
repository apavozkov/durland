package core

import (
	"durland/internal/domain"
	"durland/internal/registry"
)

type Durlandets struct {
	Race     domain.RaceType
	Nation   domain.NationType
	Stats    domain.Stats
	Location *domain.Location
	Step     int
}

func NewDurlandets(race domain.RaceType, nation domain.NationType) *Durlandets {
	return &Durlandets{
		Race:   race,
		Nation: nation,
		Stats:  domain.NewBaseStats(),
		Step:   0,
	}
}

func (d *Durlandets) CalculateActivityEffect(activity domain.ActivityType) domain.Stats {
	var baseEffect domain.Stats

	// Базовый эффект от деятельности
	switch activity {
	case domain.DoingNothing:
		baseEffect = domain.Stats{Health: -0.5, Money: -0.5, Satisfaction: -0.5}
	case domain.Zumbalit:
		slesandraCount := d.Location.GetFaunaCount(domain.Slesandra)
		baseEffect = domain.Stats{
			Health:       -1,
			Money:        2.0 * float64(slesandraCount),
			Satisfaction: -1,
		}
	case domain.Gulbonit:
		sisandraCount := d.Location.GetFaunaCount(domain.Sisandra)
		baseEffect = domain.Stats{
			Health:       -1,
			Money:        -1,
			Satisfaction: 2.0 * float64(sisandraCount),
		}
	case domain.Shlyamsat:
		chuchundraCount := d.Location.GetFaunaCount(domain.Chuchundra)
		baseEffect = domain.Stats{
			Health:       2.0 * float64(chuchundraCount),
			Money:        -1,
			Satisfaction: -1,
		}
	}

	// Создаем контекст
	ctx := &registry.EffectContext{
		Durlandets: d,
		Location:   d.Location,
		Activity:   activity,
		Step:       d.Step,
	}

	// Применяем модификаторы через реестр
	baseEffect = registry.ApplyModifiers(baseEffect, ctx, d.Nation, d.Location.Area)

	return baseEffect
}

func (d *Durlandets) PerformActivity(activity domain.ActivityType) {
	effect := d.CalculateActivityEffect(activity)
	d.Stats.ApplyChange(effect)
	d.Step++
}

func (d *Durlandets) ChangeLocation(newLocation *domain.Location) {
	d.Location = newLocation
}

func (d *Durlandets) IsAlive() bool {
	return d.Stats.IsAlive()
}
