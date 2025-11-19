package main

import (
	"durland/internal/core"
	"durland/internal/domain"
	"durland/internal/races"
	"fmt"
)

func main() {
	// Реестр автоматически инициализируется через init() функции в модификаторах
	// Добавь в начало main() для демонстрации:
	fmt.Println("=== ИНФОРМАЦИЯ О РАСАХ ===")
	for _, race := range races.GetAllRaces() {
		fmt.Printf("Раса: %s\n", race.Name)
		fmt.Printf("Народы: ")
		for _, nation := range race.AllowedNations {
			fmt.Printf("%s ", getNationName(nation))
		}
		fmt.Printf("\n---\n")
	}
	// Создаем локации
	workland := domain.NewLocation(domain.Workland, domain.Balbesburg)
	beachland := domain.NewLocation(domain.Beachland, domain.PuntaPelikana)
	pranaland := domain.NewLocation(domain.Pranaland, domain.HareKirishi)

	// Создаем дурляндцев разных народов
	mozhory := core.NewDurlandets(domain.Shlendrik, domain.Mozhory)
	mozhory.ChangeLocation(workland)

	soyevye := core.NewDurlandets(domain.Hipstik, domain.Soyevye)
	soyevye.ChangeLocation(beachland)

	drotsenty := core.NewDurlandets(domain.Skufik, domain.Drotsenty)
	drotsenty.ChangeLocation(pranaland)

	// Симуляция на 5 шагов
	fmt.Println("=== СИМУЛЯЦИЯ ДУРЛЯНДИИ ===")

	for step := 0; step < 5; step++ {
		fmt.Printf("\n--- Шаг %d ---\n", step+1)

		// Можоры зумбалят
		if mozhory.IsAlive() {
			mozhory.PerformActivity(domain.Zumbalit)
			fmt.Printf("👑 Можоры в Балбесбурге: З=%.1f 💰=%.1f 😊=%.1f\n",
				mozhory.Stats.Health, mozhory.Stats.Money, mozhory.Stats.Satisfaction)
		} else {
			fmt.Println("💀 Можоры погибли")
		}

		// Соевые гульбонят
		if soyevye.IsAlive() {
			soyevye.PerformActivity(domain.Gulbonit)
			fmt.Printf("🌱 Соевые в Пунта-пеликане: З=%.1f 💰=%.1f 😊=%.1f\n",
				soyevye.Stats.Health, soyevye.Stats.Money, soyevye.Stats.Satisfaction)
		} else {
			fmt.Println("💀 Соевые погибли")
		}

		// Дроценты шлямсают
		if drotsenty.IsAlive() {
			drotsenty.PerformActivity(domain.Shlyamsat)
			fmt.Printf("⚡ Дроценты в Харе-Кириши: З=%.1f 💰=%.1f 😊=%.1f\n",
				drotsenty.Stats.Health, drotsenty.Stats.Money, drotsenty.Stats.Satisfaction)
		} else {
			fmt.Println("💀 Дроценты погибли")
		}
	}

	fmt.Println("\n=== СИМУЛЯЦИЯ ЗАВЕРШЕНА ===")
}
func getNationName(nation domain.NationType) string {
	switch nation {
	case domain.Mozhory:
		return "Можоры"
	case domain.Nischeborody:
		return "Нищебороды"
	case domain.Soyevye:
		return "Соевые"
	case domain.Prosvetlennye:
		return "Просветленные"
	case domain.Drotsenty:
		return "Дроценты"
	case domain.Zheleznouhie:
		return "Железноухие"
	default:
		return "Неизвестно"
	}
}

// Вспомогательная функция для получения названия расы
/*func getRaceName(race domain.RaceType) string {
	switch race {
	case domain.Shlendrik:
		return "Шлендрик"
	case domain.Hipstik:
		return "Хипстик"
	case domain.Skufik:
		return "Скуфик"
	default:
		return "Неизвестно"
	}
}
*/
