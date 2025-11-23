package main

import (
	"durland/models"
	"durland/strategies"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🚀 Запуск симуляции Дурляндии...")

	// Загружаем мир
	worldDef, err := models.LoadWorldDefinition("world_definition.json")
	if err != nil {
		log.Fatalf("❌ Ошибка загрузки мира: %v", err)
	}

	races, err := models.LoadRacesDefinition("races_people.json")
	if err != nil {
		log.Fatalf("❌ Ошибка загрузки рас: %v", err)
	}

	activities, err := models.LoadActivitiesDefinition("activities.json")
	if err != nil {
		log.Fatalf("❌ Ошибка загрузки активностей: %v", err)
	}

	worldDef.Races = races
	worldDef.Activities = activities

	worldState := worldDef.ToWorldState()
	simulator := models.NewSimulator(worldState)
	strategy := &strategies.BasicStrategy{}

	// Параметры симуляции
	numDurlians := 3
	maxSteps := 20

	fmt.Printf("🎯 Симуляция: %d дурляндцев, %d шагов\n\n", numDurlians, maxSteps)

	// Статистика
	survived := 0

	// Запускаем симуляции
	for i := 0; i < numDurlians; i++ {
		durlian := models.NewDurlian(worldDef.Races, worldDef.Locations)

		fmt.Printf("🎭 ДУРЛЯНДЕЦ %d:\n", i+1)
		fmt.Printf("   🧬 %s | 👥 %s\n", durlian.Race, durlian.People)
		fmt.Printf("   📍 %s, %s\n", durlian.CurrentLocation, durlian.CurrentArea)
		fmt.Printf("   ❤️ %.1f  💰 %.1f  😊 %.1f\n\n",
			durlian.Stats.Health, durlian.Stats.Money, durlian.Stats.Satisfaction)

		fmt.Println("📝 НАЧАЛО ПОШАГОВОЙ СИМУЛЯЦИИ:")
		fmt.Println("════════════════════════════════════")

		result := simulator.RunSimulation(durlian, strategy, maxSteps)

		// Выводим пошаговую историю
		for step, history := range result.History {
			printStep(history, step+1)
		}

		fmt.Println("════════════════════════════════════")

		status := "💀 Погиб"
		if result.IsAlive {
			status = "✅ Выжил"
			survived++
		}

		fmt.Printf("🏁 РЕЗУЛЬТАТ: %s за %d шагов\n", status, result.TotalSteps)
		fmt.Printf("📊 Финальные статы: ❤️ %.1f  💰 %.1f  😊 %.1f\n\n",
			result.FinalStats.Health, result.FinalStats.Money, result.FinalStats.Satisfaction)
		fmt.Println("────────────────────────────────────")
	}

	// Финальная статистика
	fmt.Printf("📊 ИТОГИ СИМУЛЯЦИИ:\n")
	fmt.Printf("   Выжило: %d/%d (%.0f%%)\n",
		survived, numDurlians, float64(survived)/float64(numDurlians)*100)
	fmt.Println("🎉 Симуляция завершена!")
}

// Функция для красивого вывода шага
func printStep(step *models.StepHistory, stepNumber int) {
	// Эмодзи для действий
	var actionEmoji, actionText string
	switch step.Action.Type {
	case "move":
		actionEmoji = "🚶"
		actionText = fmt.Sprintf("Перемещение в %s, %s", step.Location, step.Area)
	case "activity":
		actionEmoji = "🎯"
		activityName := getRussianActivityName(step.Activity)
		actionText = fmt.Sprintf("%s в %s, %s", activityName, step.Location, step.Area)
	case "stay":
		actionEmoji = "💤"
		actionText = fmt.Sprintf("Остался в %s, %s", step.Location, step.Area)
	default:
		actionEmoji = "❓"
		actionText = "Неизвестное действие"
	}

	// Эмодзи для активности
	var activityEmoji string
	switch step.Activity {
	case "zumbalit":
		activityEmoji = "💃"
	case "gulbonit":
		activityEmoji = "🎉"
	case "shlyamsat":
		activityEmoji = "🎭"
	case "none":
		activityEmoji = "😴"
	default:
		activityEmoji = "❓"
	}

	fmt.Printf("🔄 ШАГ %d:\n", stepNumber)
	fmt.Printf("   %s %s\n", actionEmoji, actionText)

	if step.Action.Type == "activity" {
		fmt.Printf("   %s Активность: %s\n", activityEmoji, getRussianActivityName(step.Activity))
	}

	// Изменения статов
	fmt.Printf("   📈 Изменения: ")
	changes := []string{}

	if step.Effects.HealthChange != 0 {
		emoji := "🔻"
		if step.Effects.HealthChange > 0 {
			emoji = "🔺"
		}
		changes = append(changes, fmt.Sprintf("%s❤️ %+.1f", emoji, step.Effects.HealthChange))
	}

	if step.Effects.MoneyChange != 0 {
		emoji := "🔻"
		if step.Effects.MoneyChange > 0 {
			emoji = "🔺"
		}
		changes = append(changes, fmt.Sprintf("%s💰 %+.1f", emoji, step.Effects.MoneyChange))
	}

	if step.Effects.SatisfactionChange != 0 {
		emoji := "🔻"
		if step.Effects.SatisfactionChange > 0 {
			emoji = "🔺"
		}
		changes = append(changes, fmt.Sprintf("%s😊 %+.1f", emoji, step.Effects.SatisfactionChange))
	}

	if len(changes) == 0 {
		fmt.Printf("нет изменений")
	} else {
		for i, change := range changes {
			if i > 0 {
				fmt.Printf(" | ")
			}
			fmt.Printf(change)
		}
	}
	fmt.Printf("\n")

	// Текущие статы
	fmt.Printf("   📊 Статы: ❤️ %.1f | 💰 %.1f | 😊 %.1f\n\n",
		step.StatsAfter.Health, step.StatsAfter.Money, step.StatsAfter.Satisfaction)
}

// Функция для русских названий активностей
func getRussianActivityName(activity string) string {
	switch activity {
	case "zumbalit":
		return "Зумбальство"
	case "gulbonit":
		return "Гульбонство"
	case "shlyamsat":
		return "Шлямсанье"
	case "none":
		return "Отдых"
	default:
		return activity
	}
}
