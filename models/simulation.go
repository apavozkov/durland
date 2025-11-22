package models

// Состояние мира
type WorldState struct {
	Races      []Race     `json:"races"`
	Locations  []Location `json:"locations"`
	Activities []Activity `json:"activities"`
	Step       int        `json:"step"`
}

// Действие дурляндца
type Action struct {
	Type           string `json:"type"` // премещение, деятельность
	TargetLocation string `json:"target_location,omitempty"`
	TargetArea     string `json:"target_area,omitempty"`
	Activity       string `json:"activity,omitempty"`
}

// Результат применения эффектов
type EffectResult struct {
	HealthChange       float64 `json:"health_change"`
	MoneyChange        float64 `json:"money_change"`
	SatisfactionChange float64 `json:"satisfaction_change"`
}

// Результат симуляции для одного дурляндца (для статы)
type SimulationResult struct {
	DurlianID  int64          `json:"durlian_id"`
	TotalSteps int            `json:"total_steps"`
	IsAlive    bool           `json:"is_alive"`
	FinalStats Stats          `json:"final_stats"`
	Race       string         `json:"race"`
	People     string         `json:"people"`
	History    []*StepHistory `json:"history"`
}

// История одного шага
type StepHistory struct {
	Step        int          `json:"step"`
	Location    string       `json:"location"`
	Area        string       `json:"area"`
	Activity    string       `json:"activity"`
	StatsBefore Stats        `json:"stats_before"`
	StatsAfter  Stats        `json:"stats_after"`
	Effects     EffectResult `json:"effects"`
	Action      Action       `json:"action"`
}
