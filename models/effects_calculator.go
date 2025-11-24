package models

import (
	"math/rand"
)

type EffectsCalculator struct {
	stayCounts        map[int64]map[string]int
	effectRegistry    map[string]EffectHandler
	conditionRegistry map[string]ConditionChecker
}

type EffectHandler func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{})
type ConditionChecker func(durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) bool

func NewEffectsCalculator() *EffectsCalculator {
	ec := &EffectsCalculator{
		stayCounts:        make(map[int64]map[string]int),
		effectRegistry:    make(map[string]EffectHandler),
		conditionRegistry: make(map[string]ConditionChecker),
	}

	ec.registerDefaultHandlers()
	ec.registerDefaultConditions()

	return ec
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

// Базовые эффекты занятий (из JSON)
func (ec *EffectsCalculator) applyBaseActivityEffects(result *EffectResult, activity string, durlian *Durlian, worldState *WorldState) {
	// Находим активность в конфиге
	var activityConfig *Activity
	for i := range worldState.Activities {
		if worldState.Activities[i].Name == activity {
			activityConfig = &worldState.Activities[i]
			break
		}
	}

	if activityConfig == nil {
		// Дефолтная активность - ничего не делать
		activityConfig = &Activity{
			Name: "none",
			BaseEffects: []Effect{
				{
					Type:       "add_health",
					Parameters: map[string]interface{}{"value": -0.5},
				},
				{
					Type:       "add_money",
					Parameters: map[string]interface{}{"value": -0.5},
				},
				{
					Type:       "add_satisfaction",
					Parameters: map[string]interface{}{"value": -0.5},
				},
			},
		}
	}

	// Применяем все эффекты активности
	ec.applyEffects(result, activityConfig.BaseEffects, durlian, activity, worldState)
	ec.applyEffects(result, activityConfig.FaunaEffects, durlian, activity, worldState)
}

// Универсальное применение эффектов
func (ec *EffectsCalculator) applyEffects(result *EffectResult, effects []Effect, durlian *Durlian, activity string, worldState *WorldState) {
	for _, effect := range effects {
		if ec.checkConditions(effect.Conditions, durlian, activity, worldState) {
			if handler, exists := ec.effectRegistry[effect.Type]; exists {
				handler(result, durlian, activity, worldState, effect.Parameters)
			}
		}
	}
}

// Проверка условий применения эффекта
func (ec *EffectsCalculator) checkConditions(conditions []EffectCondition, durlian *Durlian, activity string, worldState *WorldState) bool {
	for _, condition := range conditions {
		if checker, exists := ec.conditionRegistry[condition.Type]; exists {
			if !checker(durlian, activity, worldState, condition.Parameters) {
				return false
			}
		}
	}
	return true
}

// Эффекты народов
func (ec *EffectsCalculator) applyPeopleEffects(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState) {
	for _, race := range worldState.Races {
		if race.Name == durlian.Race {
			for _, people := range race.Peoples {
				if people.Name == durlian.People {
					ec.applyEffects(result, people.Effects, durlian, activity, worldState)
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

	if area != nil {
		ec.applyEffects(result, area.Effects, durlian, activity, worldState)
	}
}

// Регистрация обработчиков эффектов
func (ec *EffectsCalculator) registerDefaultHandlers() {
	// Простые добавления к статам
	ec.effectRegistry["add_health"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if value, ok := params["value"].(float64); ok {
			result.HealthChange += value
		}
	}

	ec.effectRegistry["add_money"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if value, ok := params["value"].(float64); ok {
			result.MoneyChange += value
		}
	}

	ec.effectRegistry["add_satisfaction"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if value, ok := params["value"].(float64); ok {
			result.SatisfactionChange += value
		}
	}

	// Базовые модификаторы
	ec.effectRegistry["multiply_health_change"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if multiplier, ok := params["multiplier"].(float64); ok {
			result.HealthChange *= multiplier
		}
	}

	ec.effectRegistry["multiply_money_change"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if multiplier, ok := params["multiplier"].(float64); ok {
			result.MoneyChange *= multiplier
		}
	}

	ec.effectRegistry["multiply_satisfaction_change"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if multiplier, ok := params["multiplier"].(float64); ok {
			result.SatisfactionChange *= multiplier
		}
	}

	ec.effectRegistry["multiply_all_changes"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if multiplier, ok := params["multiplier"].(float64); ok {
			result.HealthChange *= multiplier
			result.MoneyChange *= multiplier
			result.SatisfactionChange *= multiplier
		}
	}

	// Эффекты фауны
	ec.effectRegistry["fauna_based_health"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		location := ec.findLocation(durlian.CurrentLocation, worldState)
		if location == nil {
			return
		}

		if amount, ok := params["amount_per_creature"].(float64); ok {
			if faunaType, ok := params["fauna_type"].(string); ok {
				var creatureCount int
				switch faunaType {
				case "slesandra":
					creatureCount = location.Fauna.Slesandra
				case "sisandra":
					creatureCount = location.Fauna.Sisandra
				case "chuchundra":
					creatureCount = location.Fauna.Chuchundra
				}
				result.HealthChange += amount * float64(creatureCount)
			}
		}
	}

	ec.effectRegistry["fauna_based_money"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		location := ec.findLocation(durlian.CurrentLocation, worldState)
		if location == nil {
			return
		}

		if amount, ok := params["amount_per_creature"].(float64); ok {
			if faunaType, ok := params["fauna_type"].(string); ok {
				var creatureCount int
				switch faunaType {
				case "slesandra":
					creatureCount = location.Fauna.Slesandra
				case "sisandra":
					creatureCount = location.Fauna.Sisandra
				case "chuchundra":
					creatureCount = location.Fauna.Chuchundra
				}
				result.MoneyChange += amount * float64(creatureCount)
			}
		}
	}

	ec.effectRegistry["fauna_based_satisfaction"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		location := ec.findLocation(durlian.CurrentLocation, worldState)
		if location == nil {
			return
		}

		if amount, ok := params["amount_per_creature"].(float64); ok {
			if faunaType, ok := params["fauna_type"].(string); ok {
				var creatureCount int
				switch faunaType {
				case "slesandra":
					creatureCount = location.Fauna.Slesandra
				case "sisandra":
					creatureCount = location.Fauna.Sisandra
				case "chuchundra":
					creatureCount = location.Fauna.Chuchundra
				}
				result.SatisfactionChange += amount * float64(creatureCount)
			}
		}
	}

	// Вероятностные эффекты
	ec.effectRegistry["chance_health_save"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if probability, ok := params["probability"].(float64); ok && rand.Float64() < probability {
			if healthSave, ok := params["health_save"].(float64); ok {
				result.HealthChange += healthSave
			}
		}
	}

	ec.effectRegistry["chance_fauna_money_loss"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		location := ec.findLocation(durlian.CurrentLocation, worldState)
		if location == nil {
			return
		}

		if probability, ok := params["probability"].(float64); ok && rand.Float64() < probability {
			if lossPerCreature, ok := params["loss_per_creature"].(float64); ok {
				if faunaType, ok := params["fauna_type"].(string); ok {
					var creatureCount int
					switch faunaType {
					case "slesandra":
						creatureCount = location.Fauna.Slesandra
					case "sisandra":
						creatureCount = location.Fauna.Sisandra
					case "chuchundra":
						creatureCount = location.Fauna.Chuchundra
					}
					result.MoneyChange -= lossPerCreature * float64(creatureCount)
				}
			}
		}
	}

	// Эффекты основанные на истории
	ec.effectRegistry["history_based_satisfaction"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if amount, ok := params["amount_per_creature"].(float64); ok {
			if faunaType, ok := params["fauna_type"].(string); ok {
				historySteps := 3
				if steps, ok := params["history_steps"].(float64); ok {
					historySteps = int(steps)
				}

				creatureCount := ec.countFaunaInHistory(durlian.History, worldState, faunaType, historySteps)
				result.SatisfactionChange += amount * float64(creatureCount)
			}
		}
	}

	// Установка конкретного значения
	ec.effectRegistry["set_stat_change"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if target, ok := params["target"].(string); ok {
			if value, ok := params["value"].(float64); ok {
				switch target {
				case "health_change":
					result.HealthChange = value
				case "money_change":
					result.MoneyChange = value
				case "satisfaction_change":
					result.SatisfactionChange = value
				}
			}
		}
	}

	// Эффекты урона от фауны
	ec.effectRegistry["chance_fauna_damage"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		location := ec.findLocation(durlian.CurrentLocation, worldState)
		if location == nil {
			return
		}

		if probability, ok := params["probability"].(float64); ok {
			if damage, ok := params["damage"].(float64); ok {
				if faunaType, ok := params["fauna_type"].(string); ok {
					var creatureCount int
					switch faunaType {
					case "slesandra":
						creatureCount = location.Fauna.Slesandra
					case "sisandra":
						creatureCount = location.Fauna.Sisandra
					case "chuchundra":
						creatureCount = location.Fauna.Chuchundra
					}

					for i := 0; i < creatureCount; i++ {
						if rand.Float64() < probability {
							result.HealthChange -= damage
						}
					}
				}
			}
		}
	}

	// Эффекты потери денег
	ec.effectRegistry["chance_money_wipe"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
		if probability, ok := params["probability"].(float64); ok && rand.Float64() < probability {
			if lossShare, ok := params["loss_share"].(float64); ok {
				result.MoneyChange -= durlian.Stats.Money * lossShare
			}
		}
	}
}

// Регистрация условий
func (ec *EffectsCalculator) registerDefaultConditions() {
	ec.conditionRegistry["activity_is"] = func(durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) bool {
		if requiredActivity, ok := params["activity"].(string); ok {
			return activity == requiredActivity
		}
		return false
	}

	ec.conditionRegistry["people_is"] = func(durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) bool {
		if requiredPeople, ok := params["people"].(string); ok {
			return durlian.People == requiredPeople
		}
		return false
	}

	ec.conditionRegistry["location_is"] = func(durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) bool {
		if requiredLocation, ok := params["location"].(string); ok {
			return durlian.CurrentLocation == requiredLocation
		}
		return false
	}

	ec.conditionRegistry["min_stay_count"] = func(durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) bool {
		if minCount, ok := params["min_count"].(float64); ok {
			return ec.getStayCount(durlian.ID, durlian.CurrentLocation) >= int(minCount)
		}
		return false
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

func (ec *EffectsCalculator) countFaunaInHistory(history []*StepHistory, worldState *WorldState, faunaType string, steps int) int {
	count := 0
	start := len(history) - steps
	if start < 0 {
		start = 0
	}

	for i := start; i < len(history); i++ {
		location := ec.findLocation(history[i].Location, worldState)
		if location != nil {
			switch faunaType {
			case "slesandra":
				count += location.Fauna.Slesandra
			case "sisandra":
				count += location.Fauna.Sisandra
			case "chuchundra":
				count += location.Fauna.Chuchundra
			}
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
