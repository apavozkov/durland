package models

import (
	"durland/utils"
	"math/rand"
	"time"
)

// Жизненные показатели дурляндца
type Stats struct {
	Health       float64 `json:"health"`
	Money        float64 `json:"money"`
	Satisfaction float64 `json:"satisfaction"`
}

// Хар-ки дурляндца
type Durlian struct {
	ID              int64     `json:"id"`
	Race            string    `json:"race"`   // Скрыто от дурляндца
	People          string    `json:"people"` // Скрыто от дурляндца
	CurrentLocation string    `json:"current_location"`
	CurrentArea     string    `json:"current_area"`
	CurrentActivity string    `json:"current_activity"`
	Stats           Stats     `json:"stats"`
	History         []string  `json:"history"`
	Steps           int       `json:"steps"`
	IsAlive         bool      `json:"is_alive"`
	KnownInfo       KnownInfo `json:"known_info"` // То, что знает дурляндец
}

// Информация, доступная дурляндцу о себе
type KnownInfo struct {
	ID              int64    `json:"id"`
	CurrentLocation string   `json:"current_location"`
	CurrentArea     string   `json:"current_area"`
	CurrentActivity string   `json:"current_activity"`
	Stats           Stats    `json:"stats"`
	History         []string `json:"history"`
	Steps           int      `json:"steps"`
	IsAlive         bool     `json:"is_alive"`
}

// Создаёт нового случайного дурляндца
func NewDurlian(races []Race, locations []Location) *Durlian {
	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Случайная раса и народ
	race := races[localRand.Intn(len(races))]
	people := race.Peoples[localRand.Intn(len(race.Peoples))]

	// Случайная локация и местность
	location := locations[localRand.Intn(len(locations))]
	area := location.Areas[localRand.Intn(len(location.Areas))]

	id := utils.GenerateID()

	durlian := &Durlian{
		ID:              id,
		Race:            race.Name,
		People:          people.Name,
		CurrentLocation: location.Name,
		CurrentArea:     area.Name,
		CurrentActivity: "none",
		Stats:           Stats{Health: 10, Money: 10, Satisfaction: 10},
		History:         []string{location.Name},
		Steps:           0,
		IsAlive:         true,
	}

	// Инициализируем известную информацию
	durlian.UpdateKnownInfo()

	return durlian
}

// Обновляет информацию, известную дурляндцу
func (d *Durlian) UpdateKnownInfo() {
	d.KnownInfo = KnownInfo{
		ID:              d.ID,
		CurrentLocation: d.CurrentLocation,
		CurrentArea:     d.CurrentArea,
		CurrentActivity: d.CurrentActivity,
		Stats:           d.Stats,
		History:         d.History,
		Steps:           d.Steps,
		IsAlive:         d.IsAlive,
	}
}

// Проверяет, погиб ли дурляндец
func (d *Durlian) IsDead() bool {
	return d.Stats.Health <= 0 || d.Stats.Money <= 0 || d.Stats.Satisfaction <= 0
}
