package server

import (
	"developers_today_test/db"
	"developers_today_test/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetCats(c *gin.Context) {
	cats, err := db.ListCats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": cats})
}

func HandleInsertCat(c *gin.Context) {
	var cat models.Cat
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind json"})
		return
	}

	if err := cat.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.InsertCat(cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}

func HandleUpdateCat(c *gin.Context) {
	name := c.Params.ByName("name")
	var update struct {
		Salary float64 `json:"salary"`
	}
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.UpdateCatSalary(name, update.Salary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}

func HandleGetCat(c *gin.Context) {
	name := c.Params.ByName("name")
	cat, err := db.GetCat(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": cat})
}

func HandleDeleteCat(c *gin.Context) {
	name := c.Params.ByName("name")
	if err := db.DeleteCat(name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}
