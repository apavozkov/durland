package main

import (
	"durland/models"
	simulator "durland/simulation"
	"fmt"
	"log"
)

func main() {
	// Загружаем конфигурацию симуляции
	config, err := simulator.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Starting simulation with %d durlian for %d steps\n",
		config.DurlianCount, config.SimulationSteps)

	// Загружаем данные мира
	races, err := models.LoadRacesDefinition("races.json")
	if err != nil {
		log.Fatalf("Failed to load races: %v", err)
	}

	activities, err := models.LoadActivitiesDefinition("activities.json")
	if err != nil {
		log.Fatalf("Failed to load activities: %v", err)
	}

	worldDef, err := models.LoadWorldDefinition("world_definition.json")
	if err != nil {
		log.Fatalf("Failed to load world definition: %v", err)
	}

	// Создаем состояние мира
	worldState := worldDef.ToWorldState()
	worldState.Races = races
	worldState.Activities = activities

	// Создаем стратегию и калькулятор эффектов
	strategy := models.BasicStrategy{}
	effectsCalculator := models.NewEffectsCalculator()

	// Запускаем симуляцию
	sim := simulator.NewSimulator(worldState, strategy, effectsCalculator, config)
	results := sim.RunSimulation()

	// Анализируем результаты
	aliveCount := 0
	var totalSteps int
	for _, result := range results {
		if result.IsAlive {
			aliveCount++
			totalSteps += result.TotalSteps
		}

		// Логируем информацию о каждом дурляндце
		fmt.Printf("Durlian %d (Race: %s, People: %s) - Alive: %t, Final Stats: Health=%.1f, Money=%.1f, Satisfaction=%.1f\n",
			result.DurlianID, result.Race, result.People, result.IsAlive,
			result.FinalStats.Health, result.FinalStats.Money, result.FinalStats.Satisfaction)
	}

	survivalRate := float64(aliveCount) / float64(len(results)) * 100
	averageSteps := float64(totalSteps) / float64(len(results))

	fmt.Printf("\n=== SIMULATION SUMMARY ===\n")
	fmt.Printf("Total durlian: %d\n", len(results))
	fmt.Printf("Survived: %d (%.1f%%)\n", aliveCount, survivalRate)
	fmt.Printf("Average steps survived: %.1f\n", averageSteps)
	fmt.Printf("==========================\n")

}
