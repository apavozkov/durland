package domain

type Location struct {
	Type  LocationType
	Area  AreaType
	Fauna map[FaunaType]int
}

func NewLocation(locType LocationType, area AreaType) *Location {
	l := &Location{
		Type:  locType,
		Area:  area,
		Fauna: make(map[FaunaType]int),
	}

	switch locType {
	case Workland:
		l.Fauna[Slesandra] = 3
		l.Fauna[Sisandra] = 1
		l.Fauna[Chuchundra] = 1
	case Beachland:
		l.Fauna[Slesandra] = 1
		l.Fauna[Sisandra] = 3
		l.Fauna[Chuchundra] = 1
	case Pranaland:
		l.Fauna[Slesandra] = 1
		l.Fauna[Sisandra] = 1
		l.Fauna[Chuchundra] = 3
	}

	return l
}

func (l *Location) GetFaunaCount(faunaType FaunaType) int {
	return l.Fauna[faunaType]
}
