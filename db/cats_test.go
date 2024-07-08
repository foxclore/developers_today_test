package db

import (
	"developers_today_test/models"
	"testing"
)

func SetTestEnvironment(t *testing.T) {
	dsn := "file:../database.db"
	err := SetHandler(dsn)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestInsertCat(t *testing.T) {
	SetTestEnvironment(t)
	err := InsertCat(models.Cat{
		Name:       "test",
		Experience: 0,
		Breed:      "test",
		Salary:     100,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUpdateCatSalary(t *testing.T) {
	SetTestEnvironment(t)
	err := UpdateCatSalary("test", 1400)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestGetCat(t *testing.T) {
	SetTestEnvironment(t)
	cat, err := GetCat("test")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(cat)
}

func TestListCats(t *testing.T) {
	SetTestEnvironment(t)
	cats, err := ListCats()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(cats)
}

func TestDeleteCat(t *testing.T) {
	SetTestEnvironment(t)
	err := DeleteCat("test")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
