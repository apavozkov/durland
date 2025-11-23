package models

import (
	"strings"
)

type BasicStrategy struct {
	
}


func  (v BasicStrategy) DecideNextAction(d Durlian) Action {

	steps := d.KnownInfo.Steps
	history := d.KnownInfo.History

	// Если только родился → разведка без риска
	if steps < 2 {
		return Action{Type: "activity", Activity: "none"}
	}

	if !hasPerformed(history, "test_zumba_done") {
		// если мы не в слесандровой локации, идём туда
		if d.KnownInfo.CurrentLocation != "workland" {
			return Action{Type: "move", TargetLocation: "workland"}
		}
		// пока не сделали 5 попыток — зумбим
		if count(history, "zumba") < 5 {
			return Action{Type: "activity", Activity: "zumbalit"}
		}
	}

	if !hasPerformed(history, "test_gulba_done") {
		if d.KnownInfo.CurrentLocation != "beachland" {
			return Action{Type: "move", TargetLocation: "beachland"}
		}
		if count(history, "gulba") < 5 {
			return Action{Type: "activity", Activity: "gulbonit"}
		}
	}

	if !hasPerformed(history, "test_shlyam_done") {
		if d.KnownInfo.CurrentLocation != "pranaland" {
			return Action{Type: "move", TargetLocation: "pranaland"}
		}
		if count(history, "shlyams") < 5 {
			return Action{Type: "activity",Activity: "shlyamsat"}
		}
	}

	return Action{Type: "activity", Activity: "none"}
}
func count(history []string, prefix string) int {
	n := 0
	for _, h := range history {
		if strings.HasPrefix(h, prefix) {
			n++
		}
	}
	return n
}

func hasPerformed(history []string, marker string) bool {
	for _, h := range history {
		if h == marker {
			return true
		}
	}
	return false
}
