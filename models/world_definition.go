package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// WorldDefinition задает описание мира: локации, фауна и доступные активности
type WorldDefinition struct {
	Races      []Race     `json:"races"`
	Locations  []Location `json:"locations"`
	Activities []Activity `json:"activities,omitempty"`
}

// ToWorldState превращает описание в состояние мира.
func (wd WorldDefinition) ToWorldState() *WorldState {
	return &WorldState{
		Races:      wd.Races,
		Locations:  wd.Locations,
		Activities: wd.Activities,
		Step:       0,
	}
}

// LoadWorldDefinition читает JSON с описанием мира
func LoadWorldDefinition(path string) (WorldDefinition, error) {
	if path == "" {
		return WorldDefinition{}, fmt.Errorf("world definition path is required")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return WorldDefinition{}, fmt.Errorf("read world definition: %w", err)
	}

	if len(data) == 0 {
		return WorldDefinition{}, fmt.Errorf("world definition file is empty: %s", path)
	}

	var def WorldDefinition
	if err := json.Unmarshal(data, &def); err != nil {
		return WorldDefinition{}, fmt.Errorf("decode world definition: %w", err)
	}

	if len(def.Locations) == 0 && len(def.Activities) == 0 {
		return WorldDefinition{}, fmt.Errorf("world definition has no locations or activities: %s", path)
	}

	return def, nil
}

// LoadRacesDefinition читает JSON с описанием рас и народов
func LoadRacesDefinition(path string) ([]Race, error) {
	if path == "" {
		return nil, fmt.Errorf("races definition path is required")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read races definition: %w", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("races definition file is empty: %s", path)
	}

	var racesDef struct {
		Races []Race `json:"races"`
	}
	if err := json.Unmarshal(data, &racesDef); err != nil {
		return nil, fmt.Errorf("decode races definition: %w", err)
	}

	if len(racesDef.Races) == 0 {
		return nil, fmt.Errorf("races definition has no races: %s", path)
	}

	return racesDef.Races, nil
}

// LoadActivitiesDefinition читает JSON с описанием активностей
func LoadActivitiesDefinition(path string) ([]Activity, error) {
	if path == "" {
		return nil, fmt.Errorf("activities definition path is required")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read activities definition: %w", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("activities definition file is empty: %s", path)
	}

	var activitiesDef struct {
		Activities []Activity `json:"activities"`
	}
	if err := json.Unmarshal(data, &activitiesDef); err != nil {
		return nil, fmt.Errorf("decode activities definition: %w", err)
	}

	if len(activitiesDef.Activities) == 0 {
		return nil, fmt.Errorf("activities definition has no activities: %s", path)
	}

	return activitiesDef.Activities, nil
}
