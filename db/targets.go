package db

import (
	"developers_today_test/models"
	"errors"
)

var (
	ErrorTargetAlreadyExists    = errors.New("target already exists")
	ErrorTargetDoesntExist      = errors.New("target doesn't exist")
	ErrorTargetAlreadyCompleted = errors.New("target already completed")
)

func InsertTarget(t models.Target) error {
	exists, err := TargetExists(t.Name)
	if err != nil {
		return err
	}
	if exists {
		return ErrorTargetAlreadyExists
	}
	H.M.Lock()
	defer H.M.Unlock()

	_, err = H.DB.Exec("INSERT INTO targets (target_id, name, completed, notes, country) "+
		"VALUES ($1, $2, $3, $4, $5)", t.TargetId, t.Name, t.Completed, t.Notes, t.Country)
	return err
}

func UpdateTargetComplete(targetName string) error {
	exists, err := TargetExists(targetName)
	if err != nil {
		return err
	}
	if !exists {
		return ErrorTargetDoesntExist
	}
	H.M.Lock()
	defer H.M.Unlock()
	_, err = H.DB.Exec("update targets set completed=1 where name=$2", targetName)
	return err
}

func UpdateTargetNotes(missionId, name, notes string) error {
	exists, err := TargetExists(name)
	if err != nil {
		return err
	}
	if !exists {
		return ErrorTargetDoesntExist
	}
	target, err := GetTarget(name)
	if err != nil {
		return err
	}
	if target.Completed {
		return ErrorTargetAlreadyCompleted
	}
	mission, err := GetMission(missionId)
	if err != nil {
		return err
	}
	if mission.Completed {
		return ErrorMissionAlreadyCompleted
	}

	H.M.Lock()
	defer H.M.Unlock()
	_, err = H.DB.Exec("update targets set notes=$1 where name=$2", notes, name)
	return err
}

func GetTarget(name string) (models.Target, error) {
	exists, err := TargetExists(name)
	if err != nil {
		return models.Target{}, err
	}
	if !exists {
		return models.Target{}, ErrorTargetDoesntExist
	}
	H.M.Lock()
	defer H.M.Unlock()
	var target models.Target
	return target, H.DB.Get(&target, "select * from targets where name=$1", name)
}

func TargetExists(name string) (bool, error) {
	H.M.Lock()
	defer H.M.Unlock()
	var data []models.Target
	err := H.DB.Select(&data, "select * from targets where name=$1", name)
	if err != nil {
		return false, err
	}
	return len(data) > 0, err
}
