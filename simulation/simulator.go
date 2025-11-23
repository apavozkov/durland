package simulator

import (
	"durland/interfaces"
	"durland/models"
	"encoding/json"
	"fmt"
	"os"
)

type SimulatorConfig struct {
	SimulationSteps int `json:"simulation_steps"`
	DurlianCount    int `json:"durlian_count"`
}

type Simulator struct {
	worldState        *models.WorldState
	strategy          interfaces.Strategy
	effectsCalculator interfaces.EffectsCalculator
	config            SimulatorConfig
}

func NewSimulator(worldState *models.WorldState, strategy interfaces.Strategy,
	effectsCalculator interfaces.EffectsCalculator, config SimulatorConfig) *Simulator {
	return &Simulator{
		worldState:        worldState,
		strategy:          strategy,
		effectsCalculator: effectsCalculator,
		config:            config,
	}
}

func (s *Simulator) RunSimulation() []*models.SimulationResult {
	results := make([]*models.SimulationResult, 0, s.config.DurlianCount)

	for i := 0; i < s.config.DurlianCount; i++ {
		// Создаем нового дурляндца
		durlian := models.NewDurlian(s.worldState.Races, s.worldState.Locations)

		result := s.simulateDurlian(durlian)
		results = append(results, result)
	}

	return results
}

func (s *Simulator) simulateDurlian(durlian *models.Durlian) *models.SimulationResult {
	result := &models.SimulationResult{
		DurlianID:  durlian.ID,
		TotalSteps: s.config.SimulationSteps,
		Race:       durlian.Race,
		People:     durlian.People,
		History:    make([]*models.StepHistory, 0),
	}

	for step := 1; step <= s.config.SimulationSteps && durlian.IsAlive; step++ {
		s.worldState.Step = step

		// Сохраняем состояние до шага
		statsBefore := durlian.Stats

		// Получаем решение от стратегии
		action := s.strategy.DecideNextAction(durlian, s.worldState)
		if action == nil {
			// Стратегия не вернула действие - ничего не делаем
			action = &models.Action{
				Type:     "activity",
				Activity: "none",
			}
		}

		// Применяем действие
		stepHistory := s.applyAction(durlian, action, step, statsBefore)
		result.History = append(result.History, stepHistory)

		// Проверяем выживаемость
		if durlian.IsDead() {
			durlian.IsAlive = false
			break
		}

		durlian.Steps++
		durlian.UpdateKnownInfo()
	}

	result.IsAlive = durlian.IsAlive
	result.FinalStats = durlian.Stats

	return result
}

func (s *Simulator) applyAction(durlian *models.Durlian, action *models.Action,
	step int, statsBefore models.Stats) *models.StepHistory {

	// Обрабатываем перемещение
	if action.Type == "move" && action.TargetLocation != "" {
		durlian.CurrentLocation = action.TargetLocation
		if action.TargetArea != "" {
			durlian.CurrentArea = action.TargetArea
		}
		durlian.History = append(durlian.History, durlian.CurrentLocation)
	}

	// Обрабатываем деятельность
	if action.Activity != "" {
		durlian.CurrentActivity = action.Activity
	}

	// Находим активность
	var activity *models.Activity
	for i := range s.worldState.Activities {
		if s.worldState.Activities[i].Name == durlian.CurrentActivity {
			activity = &s.worldState.Activities[i]
			break
		}
	}

	if activity == nil {
		// Дефолтная активность - ничего не делать
		activity = &models.Activity{Name: "none"}
	}

	// Применяем эффекты
	effectResult := s.effectsCalculator.CalculateEffects(durlian, activity, s.worldState)

	// Получаем информацию о фауне
	faunaInfo := s.effectsCalculator.GetFaunaInfo(durlian, s.worldState)

	// Проверяем на критические события
	isCritical := s.checkCriticalEvent(durlian, statsBefore, *effectResult)

	// Обновляем статистику
	durlian.Stats.Health += effectResult.HealthChange
	durlian.Stats.Money += effectResult.MoneyChange
	durlian.Stats.Satisfaction += effectResult.SatisfactionChange

	// Создаем историю шага
	return &models.StepHistory{
		Step:             step,
		Location:         durlian.CurrentLocation,
		Area:             durlian.CurrentArea,
		Activity:         durlian.CurrentActivity,
		StatsBefore:      statsBefore,
		StatsAfter:       durlian.Stats,
		Effects:          *effectResult,
		Action:           *action,
		FaunaEncountered: faunaInfo,
		IsCriticalEvent:  isCritical,
		Notes:            s.generateStepNotes(durlian, *effectResult, isCritical),
	}
	durlian.AddHistoryEntry(stepHistory)

	return stepHistory
}

// Проверяет критические события (почти смерть, резкие изменения и т.п.)
func (s *Simulator) checkCriticalEvent(durlian *models.Durlian, statsBefore models.Stats, effectResult models.EffectResult) bool {
	// Проверяем, не упал ли показатель ниже 2 (критический уровень)
	if durlian.Stats.Health <= 2 || durlian.Stats.Money <= 2 || durlian.Stats.Satisfaction <= 2 {
		return true
	}

	// Проверяем резкие изменения (более 5 единиц за шаг)
	healthChange := effectResult.HealthChange
	moneyChange := effectResult.MoneyChange
	satisfactionChange := effectResult.SatisfactionChange

	if healthChange <= -3 || healthChange >= 3 ||
		moneyChange <= -3 || moneyChange >= 3 ||
		satisfactionChange <= -3 || satisfactionChange >= 3 {
		return true
	}

	return false
}

func LoadConfig(path string) (SimulatorConfig, error) {
	var config SimulatorConfig

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	if config.SimulationSteps <= 0 {
		config.SimulationSteps = 100
	}
	if config.DurlianCount <= 0 {
		config.DurlianCount = 10
	}

	return config, nil
}

// Генерирует заметки для шага
func (s *Simulator) generateStepNotes(durlian *models.Durlian, effectResult models.EffectResult, isCritical bool) string {
	notes := ""

	if isCritical {
		notes += "Критическое событие! "

		if durlian.Stats.Health <= 2 {
			notes += "Здоровье на критически низком уровне. "
		}
		if durlian.Stats.Money <= 2 {
			notes += "Деньги на исходе. "
		}
		if durlian.Stats.Satisfaction <= 2 {
			notes += "Удовлетворенность жизнью почти нулевая. "
		}
	}

	// Добавляем информацию о значительных изменениях
	if effectResult.HealthChange != 0 || effectResult.MoneyChange != 0 || effectResult.SatisfactionChange != 0 {
		notes += fmt.Sprintf("Изменения: Здоровье(%+.1f), Деньги(%+.1f), Удовлетворенность(%+.1f)",
			effectResult.HealthChange, effectResult.MoneyChange, effectResult.SatisfactionChange)
	}

	return notes
}
