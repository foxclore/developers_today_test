package db

import (
	"developers_today_test/models"
	"errors"
	"log"
)

var (
	ErrorMissionAlreadyExists    = errors.New("mission already exists")
	ErrorMissionDoesntExist      = errors.New("mission doesn't exist")
	ErrorCatAlreadyAssigned      = errors.New("cat already assigned to mission")
	ErrorMissionAlreadyCompleted = errors.New("mission already completed")
)

func InsertMission(em models.ExportedMission) error {
	exists, err := MissionExists(em.MissionId)
	if err != nil {
		return err
	}
	if exists {
		return ErrorMissionAlreadyExists
	}
	var m = models.Mission{
		MissionId:       em.MissionId,
		AssignedCatName: em.AssignedCatName,
		Completed:       em.Completed,
	}
	for i, target := range em.Targets {
		switch i {
		case 0:
			m.Target1Name = target.Name
			break
		case 1:
			m.Target2Name = target.Name
			break
		case 2:
			m.Target3Name = target.Name
			break
		}

		if err := InsertTarget(target); err != nil {
			return err
		}
	}
	H.M.Lock()
	defer H.M.Unlock()
	_, err = H.DB.Exec("INSERT INTO missions (mission_id, target_1_name, target_2_name, "+
		"target_3_name,assigned_cat_name,completed) "+
		"VALUES ($1, $2, $3, $4, $5, $6)", m.MissionId, m.Target1Name, m.Target2Name, m.Target3Name, m.AssignedCatName, m.Completed)
	return err
}

func DeleteMission(missionId string) error {
	// Checks for existence are built into GetMission
	mission, err := GetMission(missionId)
	if err != nil {
		return err
	}
	if mission.AssignedCatName != "" {
		return ErrorCatAlreadyAssigned
	}
	H.M.Lock()
	defer H.M.Unlock()
	_, err = H.DB.Exec("delete from missions where mission_id=$1", missionId)
	return err
}

func SetMissionCompleted(missionId string) error {
	exists, err := MissionExists(missionId)
	if err != nil {
		return err
	}
	if !exists {
		return ErrorMissionDoesntExist
	}
	H.M.Lock()
	defer H.M.Unlock()
	_, err = H.DB.Exec("update missions set completed=1 where mission_id=$2", missionId)
	return err
}

func GetMission(missionId string) (models.Mission, error) {
	exists, err := MissionExists(missionId)
	if err != nil {
		return models.Mission{}, err
	}
	if !exists {
		return models.Mission{}, ErrorMissionDoesntExist
	}
	H.M.Lock()
	defer H.M.Unlock()
	var mission models.Mission
	return mission, H.DB.Get(&mission, "select * from missions where mission_id=$1", missionId)
}

func DeleteTargetFromMission(missionId, targetName string) error {
	mission, err := GetMission(missionId)
	if err != nil {
		return err
	}
	target, err := GetTarget(targetName)
	if err != nil {
		return err
	}
	if target.Completed {
		return errors.New("cannot delete target because it is completed")
	}
	H.M.Lock()
	_, err = H.DB.Exec("delete from targets where name=$1", targetName)
	if err != nil {
		return err
	}
	field := ""
	if mission.Target1Name == targetName {
		field = "target_1_name"
	} else if mission.Target2Name == targetName {
		field = "target_2_name"
	} else if mission.Target3Name == targetName {
		field = "target_3_name"
	}
	H.M.Unlock()
	return UpdateMissionField(missionId, field, "")
}

func UpdateMissionField(missionId, field, val string) error {
	exists, err := MissionExists(missionId)
	if err != nil {
		return err
	}
	if !exists {
		return ErrorMissionDoesntExist
	}
	H.M.Lock()
	defer H.M.Unlock()
	log.Println("update missions set " + field + "=" + val + " where mission_id=$1")
	_, err = H.DB.Exec("update missions set "+field+"='"+val+"' where mission_id=$1", missionId)
	return err
}

func AddTargetToMission(missionId string, target models.Target) error {
	mission, err := GetMission(missionId)
	if err != nil {
		return err
	}
	if mission.Completed {
		return errors.New("cannot add target because mission is already completed")
	}
	if mission.Target1Name != "" && mission.Target2Name != "" && mission.Target3Name != "" {
		return errors.New("mission already is at maximum capacity with targets")
	}
	err = InsertTarget(target)
	if err != nil {
		return err
	}
	field := ""
	if mission.Target1Name == "" {
		field = "target_1_name"
	} else if mission.Target2Name == "" {
		field = "target_2_name"
	} else if mission.Target3Name == "" {
		field = "target_3_name"
	}
	return UpdateMissionField(missionId, field, target.Name)
}

func UpdateMissionSetCat(catName, missionId string) error {
	exists, err := MissionExists(missionId)
	if err != nil {
		return err
	}
	if !exists {
		return ErrorMissionDoesntExist
	}
	return UpdateMissionField(missionId, "assigned_cat_name", catName)
}

func ListAllMissions() ([]models.ExportedMission, error) {
	H.M.Lock()
	var missions []models.Mission
	err := H.DB.Select(&missions, "select * from missions")
	if err != nil {
		return nil, err
	}
	H.M.Unlock()
	var res []models.ExportedMission
	for _, m := range missions {
		sp, err := GetMissionSpecial(m.MissionId)
		if err != nil {
			return nil, err
		}
		res = append(res, sp)
	}
	return res, nil
}

func GetMissionSpecial(missionId string) (models.ExportedMission, error) {
	mission, err := GetMission(missionId)
	if err != nil {
		return models.ExportedMission{}, err
	}
	var targets []models.Target
	if mission.Target1Name != "" {
		target, err := GetTarget(mission.Target1Name)
		if err != nil {
			return models.ExportedMission{}, err
		}
		targets = append(targets, target)
	}
	if mission.Target2Name != "" {
		target, err := GetTarget(mission.Target2Name)
		if err != nil {
			return models.ExportedMission{}, err
		}
		targets = append(targets, target)
	}
	if mission.Target3Name != "" {
		target, err := GetTarget(mission.Target3Name)
		if err != nil {
			return models.ExportedMission{}, err
		}
		targets = append(targets, target)
	}
	return models.ExportedMission{
		MissionId:       missionId,
		Targets:         targets,
		Completed:       mission.Completed,
		AssignedCatName: mission.AssignedCatName,
	}, nil
}

func MissionExists(missionId string) (bool, error) {
	H.M.Lock()
	defer H.M.Unlock()
	var data []models.Mission
	err := H.DB.Select(&data, "select * from missions where mission_id=$1", missionId)
	if err != nil {
		return false, err
	}
	return len(data) > 0, err
}
