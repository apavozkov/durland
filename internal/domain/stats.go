package domain

type Stats struct {
	Health       float64
	Money        float64
	Satisfaction float64
}

func (s *Stats) IsAlive() bool {
	return s.Health > 0 && s.Money > 0 && s.Satisfaction > 0
}

func (s *Stats) ApplyChange(delta Stats) {
	s.Health += delta.Health
	s.Money += delta.Money
	s.Satisfaction += delta.Satisfaction
}

func NewBaseStats() Stats {
	return Stats{
		Health:       10,
		Money:        10,
		Satisfaction: 10,
	}
}
