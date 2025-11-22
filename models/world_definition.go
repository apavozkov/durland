package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// WorldDefinition задает описание мира: локации, фауна и доступные активности
type WorldDefinition struct {
	Locations  []Location `json:"locations"`
	Activities []Activity `json:"activities,omitempty"`
}

// ToWorldState превращает описание в состояние мира.
func (wd WorldDefinition) ToWorldState() *WorldState {
	return &WorldState{
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
