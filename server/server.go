package server

import "github.com/gin-gonic/gin"

func SetServer(r *gin.Engine) {
	r.GET("/cats", HandleGetCats)
	r.POST("/cats", HandleInsertCat)
	r.PUT("/cats/:name", HandleUpdateCat)
	r.GET("/cats/:name", HandleGetCat)
	r.DELETE("/cats/:name", HandleDeleteCat)

	r.GET("/missions", HandleListAllMissions)
	r.POST("/missions", HandleCreateMission)
	r.DELETE("/missions/:mission", HandleDeleteMission)
	r.PUT("/missions/:mission", HandleUpdateMission)
	r.POST("/missions/:mission/targets", HandleAddTarget)
	r.PUT("/missions/:mission/notes", HandleUpdateTargetNotes)
	r.DELETE("/missions/", HandleDeleteMissionTarget)
	r.PUT("/missions/targets/:target", HandleUpdateTargetComplete)
	r.PUT("/missions/:mission/cat", HandleAddCat)
	r.GET("/missions/:mission", HandleGetMission)
}
