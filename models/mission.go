package models

import (
	"developers_today_test/utils"
	"errors"
)

type Mission struct {
	MissionId       string `json:"mission_id" db:"mission_id"`
	Target1Name     string `json:"target_1_name" db:"target_1_name"`
	Target2Name     string `json:"target_2_name" db:"target_2_name"`
	Target3Name     string `json:"target_3_name" db:"target_3_name"`
	AssignedCatName string `json:"assigned_cat_name" db:"assigned_cat_name"`
	Completed       bool   `json:"completed" db:"completed"`
}

type ExportedMission struct {
	MissionId       string   `json:"mission_id,omitempty"`
	Targets         []Target `json:"targets"`
	Completed       bool     `json:"completed"`
	AssignedCatName string   `json:"assigned_cat_name,omitempty"`
}

func (em *ExportedMission) SetId() {
	if em.MissionId == "" {
		em.MissionId = utils.RandomString()
	}
	for i := range em.Targets {
		if em.Targets[i].TargetId == "" {
			em.Targets[i].TargetId = utils.RandomString(7)
		}
	}
}

func (em *ExportedMission) Verify() error {
	if len(em.Targets) > 3 {
		return errors.New("too many targets")
	}
	if len(em.Targets) == 0 {
		return errors.New("set at least one target")
	}
	return nil
}
