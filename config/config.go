package config

import (
	"durland/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Races      []models.Race     `json:"races"`
	Locations  []models.Location `json:"locations"`
	Activities []models.Activity `json:"activities"`
}

var GlobalConfig *Config

// Загружает конфигурацию из JSON
func LoadConfig(configPath string) error {
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return fmt.Errorf("не удалось получить абсолютный путь: %v", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл конфигурации: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("не удалось спарсить JSON: %v", err)
	}

	// Валидация конфигурации
	if err := validateConfig(&config); err != nil {
		return fmt.Errorf("не удалось провалидировать конфигурацию: %v", err)
	}

	GlobalConfig = &config
	log.Printf("Конфигурация загружена из: %s", absPath)
	return nil
}

// Возвращает глобальную конфигурацию
func GetConfig() *Config {
	if GlobalConfig == nil {
		log.Fatal("Конфигурация не загружена")
	}
	return GlobalConfig
}

// Проверяет корректность конфигурации
func validateConfig(config *Config) error {
	if len(config.Races) == 0 {
		return fmt.Errorf("не заданы расы")
	}
	if len(config.Locations) == 0 {
		return fmt.Errorf("не заданы локации")
	}
	if len(config.Activities) == 0 {
		return fmt.Errorf("не заданы занятия")
	}

	// Проверка уникальности имен
	if err := checkUniqueNames(config); err != nil {
		return err
	}

	return nil
}

func checkUniqueNames(config *Config) error {
	raceNames := make(map[string]bool)
	for _, race := range config.Races {
		if raceNames[race.Name] {
			return fmt.Errorf("дублирующееся имя расы: %s", race.Name)
		}
		raceNames[race.Name] = true
	}

	locationNames := make(map[string]bool)
	for _, location := range config.Locations {
		if locationNames[location.Name] {
			return fmt.Errorf("дублирующееся имя локации: %s", location.Name)
		}
		locationNames[location.Name] = true
	}

	activityNames := make(map[string]bool)
	for _, activity := range config.Activities {
		if activityNames[activity.Name] {
			return fmt.Errorf("дублирующееся имя занятия: %s", activity.Name)
		}
		activityNames[activity.Name] = true
	}

	return nil
}
