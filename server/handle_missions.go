package server

import (
	"developers_today_test/db"
	"developers_today_test/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleCreateMission(c *gin.Context) {
	var em models.ExportedMission
	if err := c.BindJSON(&em); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	em.SetId()
	err := em.Verify()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = db.InsertMission(em)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}

func HandleDeleteMission(c *gin.Context) {
	missionId := c.Params.ByName("mission")
	err := db.DeleteMission(missionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}

func HandleUpdateMission(c *gin.Context) {
	missionId := c.Params.ByName("mission")
	err := db.SetMissionCompleted(missionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}

func HandleUpdateTargetComplete(c *gin.Context) {
	targetName := c.Params.ByName("target")
	err := db.UpdateTargetComplete(targetName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}

func HandleUpdateTargetNotes(c *gin.Context) {
	var targetNotes struct {
		MissionId string `json:"mission_id"`
		Notes     string `json:"notes"`
	}
	targetName := c.Params.ByName("mission") // Little life-hack to get param, which is actually a target name
	if err := c.BindJSON(&targetNotes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.UpdateTargetNotes(targetNotes.MissionId, targetName, targetNotes.Notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}

func HandleDeleteMissionTarget(c *gin.Context) {
	var request struct {
		MissionId  string `json:"mission_id"`
		TargetName string `json:"target_name"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := db.DeleteTargetFromMission(request.MissionId, request.TargetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}

func HandleAddTarget(c *gin.Context) {
	missionId := c.Params.ByName("mission")
	var t models.Target
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := db.AddTargetToMission(missionId, t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}

func HandleAddCat(c *gin.Context) {
	var request struct {
		CatName string `json:"cat_name"`
	}
	missionId := c.Params.ByName("mission")
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.UpdateMissionSetCat(request.CatName, missionId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}

func HandleListAllMissions(c *gin.Context) {
	missions, err := db.ListAllMissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": missions})
}

func HandleGetMission(c *gin.Context) {
	missionId := c.Params.ByName("mission")
	mission, err := db.GetMissionSpecial(missionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": mission})
}
