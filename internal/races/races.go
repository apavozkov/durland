package races

import "durland/internal/domain"

// RaceInfo содержит информацию о расе
type RaceInfo struct {
	Type           domain.RaceType
	Name           string
	AllowedNations []domain.NationType
}

// GetRaceInfo возвращает информацию о расе
func GetRaceInfo(race domain.RaceType) RaceInfo {
	switch race {
	case domain.Shlendrik:
		return RaceInfo{
			Type: domain.Shlendrik,
			Name: "Шлендрик",
			AllowedNations: []domain.NationType{
				domain.Mozhory,
				domain.Nischeborody,
			},
		}
	case domain.Hipstik:
		return RaceInfo{
			Type: domain.Hipstik,
			Name: "Хипстик",
			AllowedNations: []domain.NationType{
				domain.Soyevye,
				domain.Prosvetlennye,
			},
		}
	case domain.Skufik:
		return RaceInfo{
			Type: domain.Skufik,
			Name: "Скуфик",
			AllowedNations: []domain.NationType{
				domain.Drotsenty,
				domain.Zheleznouhie,
			},
		}
	default:
		return RaceInfo{}
	}
}

// IsNationAllowed проверяет, может ли народ принадлежать расе
func IsNationAllowed(race domain.RaceType, nation domain.NationType) bool {
	raceInfo := GetRaceInfo(race)
	for _, allowedNation := range raceInfo.AllowedNations {
		if allowedNation == nation {
			return true
		}
	}
	return false
}

// GetRaceName возвращает название расы
func GetRaceName(race domain.RaceType) string {
	return GetRaceInfo(race).Name
}

// GetAllRaces возвращает все доступные расы
func GetAllRaces() []RaceInfo {
	return []RaceInfo{
		GetRaceInfo(domain.Shlendrik),
		GetRaceInfo(domain.Hipstik),
		GetRaceInfo(domain.Skufik),
	}
}

// GetNationsForRace возвращает все народы для указанной расы
func GetNationsForRace(race domain.RaceType) []domain.NationType {
	return GetRaceInfo(race).AllowedNations
}
