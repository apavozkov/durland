package domain

// Базовые типы и константы
type RaceType int
type NationType int
type LocationType int
type AreaType int
type ActivityType int
type FaunaType int

const (
	Shlendrik RaceType = iota
	Hipstik
	Skufik
)

const (
	Mozhory NationType = iota
	Nischeborody
	Soyevye
	Prosvetlennye
	Drotsenty
	Zheleznouhie
)

const (
	Workland LocationType = iota
	Beachland
	Pranaland
)

const (
	Balbesburg AreaType = iota
	Dolbesburg
	Kuramariby
	PuntaPelikana
	Shrinavas
	HareKirishi
)

const (
	DoingNothing ActivityType = iota
	Zumbalit
	Gulbonit
	Shlyamsat
)

const (
	Slesandra FaunaType = iota
	Sisandra
	Chuchundra
)

// String методы для удобства отладки
func (r RaceType) String() string {
	switch r {
	case Shlendrik:
		return "Шлендрик"
	case Hipstik:
		return "Хипстик"
	case Skufik:
		return "Скуфик"
	default:
		return "Неизвестно"
	}
}

func (n NationType) String() string {
	switch n {
	case Mozhory:
		return "Можоры"
	case Nischeborody:
		return "Нищебороды"
	case Soyevye:
		return "Соевые"
	case Prosvetlennye:
		return "Просветленные"
	case Drotsenty:
		return "Дроценты"
	case Zheleznouhie:
		return "Железноухие"
	default:
		return "Неизвестно"
	}
}

func (a AreaType) String() string {
	switch a {
	case Balbesburg:
		return "Балбесбург"
	case Dolbesburg:
		return "Долбесбург"
	case Kuramariby:
		return "Курамарибы"
	case PuntaPelikana:
		return "Пунта-пеликана"
	case Shrinavas:
		return "Шринавас"
	case HareKirishi:
		return "Харе-Кириши"
	default:
		return "Неизвестно"
	}
}
