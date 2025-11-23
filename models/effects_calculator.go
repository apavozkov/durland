package models

import (
	"math/rand"
)

type EffectsCalculator struct {
	stayCounts map[int64]map[string]int
}

func NewEffectsCalculator() *EffectsCalculator {
	return &EffectsCalculator{
		stayCounts: make(map[int64]map[string]int),
	}
}

// Основной метод расчета эффектов
func (ec *EffectsCalculator) CalculateEffects(durlian *Durlian, activity *Activity, worldState *WorldState) *EffectResult {
	result := &EffectResult{}

	// Обновляем счетчик пребывания
	ec.updateStayCount(durlian)

	// Базовые эффекты от деятельности
	ec.applyBaseActivityEffects(result, activity.Name, durlian, worldState)

	// Эффекты от народности
	ec.applyPeopleEffects(result, durlian, activity.Name, worldState)

	// Эффекты от местности
	ec.applyAreaEffects(result, durlian, activity.Name, worldState)

	return result
}

// Базовые эффекты занятий
func (ec *EffectsCalculator) applyBaseActivityEffects(result *EffectResult, activity string, durlian *Durlian, worldState *WorldState) {
	location := ec.findLocation(durlian.CurrentLocation, worldState)
	if location == nil {
		return
	}

	switch activity {
	case "zumbalit": // Зумбалить
		result.HealthChange -= 1.0
		result.SatisfactionChange -= 1.0
		result.MoneyChange += 2.0 * float64(location.Fauna.Slesandra)

	case "gulbonit": // Гульбонить
		result.HealthChange -= 1.0
		result.MoneyChange -= 1.0
		result.SatisfactionChange += 2.0 * float64(location.Fauna.Sisandra)

	case "shlyamsat": // Шлямсать
		result.MoneyChange -= 1.0
		result.SatisfactionChange -= 1.0
		result.HealthChange += 2.0 * float64(location.Fauna.Chuchundra)

	case "none": // Ничего не делать
		result.HealthChange -= 0.5
		result.MoneyChange -= 0.5
		result.SatisfactionChange -= 0.5
	}
}

// Эффекты народов
func (ec *EffectsCalculator) applyPeopleEffects(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState) {
	for _, race := range worldState.Races {
		if race.Name == durlian.Race {
			for _, people := range race.Peoples {
				if people.Name == durlian.People {
					for _, effect := range people.Effects {
						ec.applyPeopleEffect(result, effect, durlian, activity, worldState)
					}
					break
				}
			}
			break
		}
	}
}

// Эффекты местностей
func (ec *EffectsCalculator) applyAreaEffects(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState) {
	location := ec.findLocation(durlian.CurrentLocation, worldState)
	if location == nil {
		return
	}

	var area *Area
	for _, a := range location.Areas {
		if a.Name == durlian.CurrentArea {
			area = &a
			break
		}
	}

	if area == nil {
		return
	}

	for _, effect := range area.Effects {
		ec.applyAreaEffect(result, effect, durlian, activity, worldState)
	}
}

// Применение конкретных эффектов народов
func (ec *EffectsCalculator) applyPeopleEffect(result *EffectResult, effect Effect, durlian *Durlian, activity string, worldState *WorldState) {
	location := ec.findLocation(durlian.CurrentLocation, worldState)

	switch effect.Type {

	// Шлендрики-Можоры
	case EffectMozhoryGulbonitMoneyIncrease:
		if activity == "gulbonit" {
			multiplier := effect.Parameters["multiplier"].(float64)
			result.MoneyChange *= multiplier
		}

	case EffectMozhoryZumbalitHealthSave:
		if activity == "zumbalit" && rand.Float64() < effect.Parameters["probability"].(float64) {
			result.HealthChange = 0 // Не тратим здоровье
		}

	// Шлендрики-Нищебороды
	case EffectNishcheboryGulbonitMoneyDecrease:
		if activity == "gulbonit" {
			multiplier := effect.Parameters["multiplier"].(float64)
			result.MoneyChange *= multiplier
		}

	case EffectNishcheboryGulbonitHealthIncrease:
		if activity == "gulbonit" {
			multiplier := effect.Parameters["multiplier"].(float64)
			result.HealthChange *= multiplier
		}

	// Хипстики-Соевые
	case EffectSoyevyeZumbalitHealthPenalty:
		if activity == "zumbalit" && location != nil {
			penalty := effect.Parameters["penalty_per_chuchundra"].(float64)
			result.HealthChange -= penalty * float64(location.Fauna.Chuchundra)
		}

	// Хипстики-Просветленные
	case EffectProsvetlennyeShlyamsatHistoryBonus:
		if activity == "shlyamsat" {
			bonus := effect.Parameters["satisfaction_per_sisandra"].(float64)
			historySteps := int(effect.Parameters["history_steps"].(float64))
			sisandraCount := ec.countSisandraInHistory(durlian.History, worldState, historySteps)
			result.SatisfactionChange += bonus * float64(sisandraCount)
		}

	// Скуфики-Дроценты
	case EffectDrotsentyGulbonitEfficiencyDecrease:
		if activity == "gulbonit" {
			multiplier := effect.Parameters["multiplier"].(float64)
			result.HealthChange *= multiplier
			result.MoneyChange *= multiplier
			result.SatisfactionChange *= multiplier
		}

	// Скуфики-Железноухие
	case EffectZheleznoukhiZumbalitSatisfactionImmunity:
		if activity == "zumbalit" {
			result.SatisfactionChange = 0
		}

	case EffectZheleznoukhiZumbalitMoneyLossChance:
		if activity == "zumbalit" && rand.Float64() < effect.Parameters["probability"].(float64) {
			if location != nil {
				result.MoneyChange -= 2.0 * float64(location.Fauna.Slesandra)
			}
		}
	}
}

// Применение эффектов местностей
func (ec *EffectsCalculator) applyAreaEffect(result *EffectResult, effect Effect, durlian *Durlian, activity string, worldState *WorldState) {
	location := ec.findLocation(durlian.CurrentLocation, worldState)
	if location == nil {
		return
	}

	stayCount := ec.getStayCount(durlian.ID, durlian.CurrentLocation)

	switch effect.Type {

	// Балбесбург
	case EffectBalbesburgSlesandraDamageChance:
		probability := effect.Parameters["probability"].(float64)
		damage := effect.Parameters["damage"].(float64)
		for i := 0; i < location.Fauna.Slesandra; i++ {
			if rand.Float64() < probability {
				result.HealthChange -= damage
			}
		}

	// Долбесбург
	case EffectDolbesburgSlesandraProductivityBonus:
		if activity == "zumbalit" {
			bonus := effect.Parameters["bonus_percent"].(float64)
			result.MoneyChange *= (1 + bonus)
		}

	case EffectDolbesburgSatisfactionCostMultiplier:
		multiplier := effect.Parameters["multiplier"].(float64)
		result.SatisfactionChange *= multiplier

	// Курамарибы
	case EffectKuramaribySisandraFatigueProbability:
		if stayCount >= 2 && activity == "gulbonit" {
			probability := effect.Parameters["probability"].(float64)
			workingSisandra := 0
			for i := 0; i < location.Fauna.Sisandra; i++ {
				if rand.Float64() >= probability {
					workingSisandra++
				}
			}
			// Пересчитываем удовлетворенность только от работающих сисяндр
			if activity == "gulbonit" {
				baseSatisfaction := (result.SatisfactionChange + 1.0) / 2.0 // Восстанавливаем базовое значение
				result.SatisfactionChange = baseSatisfaction*float64(workingSisandra) - 1.0
			}
		}

	// Пунта-пеликана
	case EffectPuntaPelikanaSisandraProductivityBonus:
		if stayCount >= 2 && activity == "gulbonit" {
			bonus := effect.Parameters["bonus_percent"].(float64)
			result.SatisfactionChange *= (1 + bonus)
		}

	case EffectPuntaPelikanaMoneyWipeProbability:
		if stayCount >= 2 && rand.Float64() < effect.Parameters["probability"].(float64) {
			lossShare := effect.Parameters["loss_share"].(float64)
			result.MoneyChange -= durlian.Stats.Money * lossShare
		}

	// Шринавас
	case EffectShrinavasChuchundraProductivityBonus:
		if activity == "shlyamsat" {
			bonus := effect.Parameters["bonus_percent"].(float64)
			result.HealthChange *= (1 + bonus)
		}

	// Харе-Кириши
	case EffectHareKirishiDrotsentyHealthPenalty:
		if durlian.People == "Дроценты" {
			penalty := effect.Parameters["extra_health_percent"].(float64)
			result.HealthChange -= durlian.Stats.Health * penalty
		}
	}
}

// Вспомогательные методы
func (ec *EffectsCalculator) findLocation(locationName string, worldState *WorldState) *Location {
	for _, loc := range worldState.Locations {
		if loc.Name == locationName {
			return &loc
		}
	}
	return nil
}

func (ec *EffectsCalculator) countSisandraInHistory(history []*StepHistory, worldState *WorldState, steps int) int {
	count := 0
	start := len(history) - steps
	if start < 0 {
		start = 0
	}

	for i := start; i < len(history); i++ {
		location := ec.findLocation(history[i].Location, worldState)
		if location != nil {
			count += location.Fauna.Sisandra
		}
	}
	return count
}

func (ec *EffectsCalculator) updateStayCount(durlian *Durlian) {
	if ec.stayCounts[durlian.ID] == nil {
		ec.stayCounts[durlian.ID] = make(map[string]int)
	}
	ec.stayCounts[durlian.ID][durlian.CurrentLocation]++
}

func (ec *EffectsCalculator) getStayCount(durlianID int64, location string) int {
	if ec.stayCounts[durlianID] == nil {
		return 0
	}
	return ec.stayCounts[durlianID][location]
}

func (ec *EffectsCalculator) GetFaunaInfo(durlian *Durlian, worldState *WorldState) FaunaInfo {
	location := ec.findLocation(durlian.CurrentLocation, worldState)
	if location == nil {
		return FaunaInfo{}
	}

	return FaunaInfo{
		SlesandraCount:  location.Fauna.Slesandra,
		SisandraCount:   location.Fauna.Sisandra,
		ChuchundraCount: location.Fauna.Chuchundra,
	}
}
