package strategy

import (
	"durland/models"
)

type BasicStrategy struct {
}

// func (v BasicStrategy) DecideNextAction(d Durlian) Action {
func (v BasicStrategy) DecideNextAction(d *models.Durlian, worldState *models.WorldState) *models.Action {

	steps := d.KnownInfo.Steps
	history := d.KnownInfo.History

	// Если только родился → разведка без риска
	if steps < 2 {
		return &models.Action{Type: "activity", Activity: "none"}
	}

	if !hasPerformed(history, "test_zumba_done") {
		// если мы не в слесандровой локации, идём туда
		if d.KnownInfo.CurrentLocation != "workland" {
			return &models.Action{Type: "move", TargetLocation: "workland"}
		}
		// пока не сделали 5 попыток — зумбим
		if count(history, "zumba") < 5 {
			return &models.Action{Type: "activity", Activity: "zumbalit"}
		}
	}

	if !hasPerformed(history, "test_gulba_done") {
		if d.KnownInfo.CurrentLocation != "beachland" {
			return &models.Action{Type: "move", TargetLocation: "beachland"}
		}
		if count(history, "gulba") < 5 {
			return &models.Action{Type: "activity", Activity: "gulbonit"}
		}
	}

	if !hasPerformed(history, "test_shlyam_done") {
		if d.KnownInfo.CurrentLocation != "pranaland" {
			return &models.Action{Type: "move", TargetLocation: "pranaland"}
		}
		if count(history, "shlyams") < 5 {
			return &models.Action{Type: "activity", Activity: "shlyamsat"}
		}
	}

	return &models.Action{Type: "activity", Activity: "none"}
}
func count(history []*models.StepHistory, activity string) int {
	n := 0
	for _, h := range history {
		if h.Activity == activity {
			n++
		}
	}
	return n
}

func hasPerformed(history []*models.StepHistory, marker string) bool {
	for _, h := range history {
		if h.Notes == marker {
			return true
		}
	}
	return false
}
