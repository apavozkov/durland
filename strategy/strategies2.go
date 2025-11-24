package strategy

import (
	"durland/models"
	"math/rand"
	"time"
	// "fmt"
	// "os"
)

// RandomStrategy - простейшая стратегия-заглушка для тестирования
type RandomStrategy struct {
	rand *rand.Rand
}

func NewRandomStrategy() *RandomStrategy {
	return &RandomStrategy{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *RandomStrategy) DecideNextAction(d *models.Durlian, ws *models.WorldState) *models.Action {

	
	if !d.IsAlive {
		return &models.Action{
				Type:     "activity",
				Activity: "none",
			}
	}
	// Параметр изменяется от 1 до 5, выше его делать не имеет смысла, потому что функция весов симметрична относительно центральной оценки (5/10)
	critical_threshold := 3.0


	multiplier := 1 - critical_threshold/10
	critical_residual_weight :=(1 - multiplier) / 2
	if (multiplier + critical_residual_weight*2 !=1.0){
		panic("Sum of weights should be 1")
	}

	// Веса активностей распределены равномерно
	weights:=[]float64{0.3,0.3,0.3}

	// Если статистика опустилась ниже уровня critical_threshold, мы даем больше веса тому параметру, который больше всего в этом нуждается.
	// То есть, если здоровье 2/10, то нам надо на следующем шаге максимизировать здоровье, поэтому этой метрике мы будем давать пропорцонально больший вес.

	if (d.Stats.Health < critical_threshold){
		weights = []float64{multiplier,critical_residual_weight,critical_residual_weight}
	}
	if (d.Stats.Money < critical_threshold){
		weights = []float64{critical_residual_weight,multiplier,critical_residual_weight}
	}
	if (d.Stats.Satisfaction < critical_threshold){
		weights = []float64{critical_residual_weight,critical_residual_weight,multiplier}
	}


	actions := collectPossibleActions(d, ws)
	
	
	if len(actions) == 0 {
		return &models.Action{
				Type:     "activity",
				Activity: "none",
			}
	}
	effectsCalculator := models.NewEffectsCalculator()

	previousEffect := 0.0
	bestIter := 0
	// находим действие с максимальной взвешенной оценкой.

	for i, action := range actions {
		
		var activity *models.Activity
		
		for i := range ws.Activities {
	
			if ws.Activities[i].Name == action.Activity {
				activity = &ws.Activities[i]
				break
			}
		}
		
		if activity == nil {
			activity = &models.Activity{Name: "none"}
		}

		effect := effectsCalculator.CalculateEffects(d , activity, ws)
		if (i == 0){
			previousEffect = weights[0] * effect.HealthChange + weights[1] * effect.MoneyChange + weights[2] * effect.SatisfactionChange
		}else
		{
			currentEffect:= weights[0]* effect.HealthChange + weights[1] * effect.MoneyChange +weights[2] * effect.SatisfactionChange
			if (currentEffect > previousEffect){
				bestIter = i
				previousEffect = currentEffect
			}
		}
		
	}
	
	return actions[bestIter]



	// // Случайно выбираем тип действия: 70% - деятельность, 30% - перемещение
	// actionType := r.rand.Intn(100)

	// if actionType < 70 {
	// 	// Деятельность
	// 	activities := []string{"zumbalit", "gulbonit", "shlyamsat", "none"}
	// 	activity := activities[r.rand.Intn(len(activities))]

	// 	return &models.Action{
	// 		Type:     "activity",
	// 		Activity: activity,
	// 	}
	// } else {
	// 	// Перемещение
	// 	if len(worldState.Locations) == 0 {
	// 		// Если нет локаций, остаемся на месте
	// 		return &models.Action{
	// 			Type:     "activity",
	// 			Activity: "none",
	// 		}
	// 	}

	// 	// Выбираем случайную локацию
	// 	location := worldState.Locations[r.rand.Intn(len(worldState.Locations))]

	// 	// Выбираем случайную область в локации
	// 	var area string
	// 	if len(location.Areas) > 0 {
	// 		area = location.Areas[r.rand.Intn(len(location.Areas))].Name
	// 	} else {
	// 		area = ""
	// 	}

	// 	return &models.Action{
	// 		Type:           "move",
	// 		TargetLocation: location.Name,
	// 		TargetArea:     area,
	// 	}
	// }
}


func collectPossibleActions(d *models.Durlian, ws *models.WorldState) []*models.Action {
	var actions []*models.Action
	for l := range ws.Locations {
		if (ws.Locations[l].Name != d.CurrentLocation){
			actions = append(actions, &models.Action{
				Type:           "move",
				TargetLocation: ws.Locations[l].Name,
				Activity: "none",
			})

			}else{
				area := ""
				for _, s := range ws.Locations[l].Areas {
					if d.CurrentArea != s.Name {
						area = s.Name
					}
				}
				
				actions = append(actions, &models.Action{
				Type:           "move",
				TargetLocation: d.CurrentLocation,
				TargetArea:     area,
				Activity: "none",
			})
			}

		}
	
	activities := []string{"zumbalit", "gulbonit", "shlyamsat", "none"}

	for _, a := range  activities{

			actions = append(actions, &models.Action{
				Type:           "activity",
				
				Activity: a,
			})
		
	}

	
	return actions
}
