package models

// Раса
type Race struct {
	Name    string   `json:"name"`
	Peoples []People `json:"peoples"`
}

// Народность в расе
type People struct {
	Name    string   `json:"name"`
	Effects []Effect `json:"effects"`
}

type Effect struct {
	Type       string                 `json:"type"`
	Parameters map[string]interface{} `json:"parameters"`
	Conditions []EffectCondition      `json:"conditions,omitempty"`
}

// Условие применения эффекта
type EffectCondition struct {
	Type       string                 `json:"type"`
	Parameters map[string]interface{} `json:"parameters"`
}

// Локация
type Location struct {
	Name  string `json:"name"`
	Areas []Area `json:"areas"`
	Fauna Fauna  `json:"fauna"`
}

// Местность в локации
type Area struct {
	Name    string   `json:"name"`
	Effects []Effect `json:"effects"`
}

// Фауна в локации
type Fauna struct {
	Slesandra  int `json:"slesandra"`
	Sisandra   int `json:"sisandra"`
	Chuchundra int `json:"chuchundra"`
}

// Деятельность
type Activity struct {
	Name         string   `json:"name"`
	BaseEffects  []Effect `json:"base_effects"`
	FaunaEffects []Effect `json:"fauna_effects"`
}

// Результат действия
type Gain struct {
	Type   string  `json:"type"`   // Стата
	Fauna  string  `json:"fauna"`  // Животинка
	Amount float64 `json:"amount"` // Награда за каждую животинку
}
