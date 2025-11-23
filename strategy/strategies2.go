package strategy

import (
	"durland/models"
	"math/rand"
	"time"
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

func (r *RandomStrategy) DecideNextAction(durlian *models.Durlian, worldState *models.WorldState) *models.Action {
	// Случайно выбираем тип действия: 70% - деятельность, 30% - перемещение
	actionType := r.rand.Intn(100)

	if actionType < 70 {
		// Деятельность
		activities := []string{"zumbalit", "gulbonit", "shlyamsat", "none"}
		activity := activities[r.rand.Intn(len(activities))]

		return &models.Action{
			Type:     "activity",
			Activity: activity,
		}
	} else {
		// Перемещение
		if len(worldState.Locations) == 0 {
			// Если нет локаций, остаемся на месте
			return &models.Action{
				Type:     "activity",
				Activity: "none",
			}
		}

		// Выбираем случайную локацию
		location := worldState.Locations[r.rand.Intn(len(worldState.Locations))]

		// Выбираем случайную область в локации
		var area string
		if len(location.Areas) > 0 {
			area = location.Areas[r.rand.Intn(len(location.Areas))].Name
		} else {
			area = ""
		}

		return &models.Action{
			Type:           "move",
			TargetLocation: location.Name,
			TargetArea:     area,
		}
	}
}
