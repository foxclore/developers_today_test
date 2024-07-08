package db

import (
	"developers_today_test/models"
	"errors"
)

var (
	ErrorCatDoesntExist   = errors.New("cat doesn't exist")
	ErrorCatAlreadyExists = errors.New("cat already exists")
)

func InsertCat(cat models.Cat) error {
	exists, err := CatExists(cat.Name)
	if err != nil {
		return err
	}
	if exists {
		return ErrorCatAlreadyExists
	}
	H.M.Lock()
	defer H.M.Unlock()
	_, err = H.DB.Exec("INSERT INTO cats (name, experience, breed, salary) "+
		"VALUES ($1, $2, $3, $4)", cat.Name, cat.Experience, cat.Breed, cat.Salary)
	return err
}

func DeleteCat(name string) error {
	exists, err := CatExists(name)
	if err != nil {
		return err
	}
	if !exists {
		return ErrorCatDoesntExist
	}
	H.M.Lock()
	defer H.M.Unlock()
	_, err = H.DB.Exec("delete from cats where name=$1", name)
	return err
}

func UpdateCatSalary(name string, salary float64) error {
	exists, err := CatExists(name)
	if err != nil {
		return err
	}
	if !exists {
		return ErrorCatDoesntExist
	}
	H.M.Lock()
	defer H.M.Unlock()
	_, err = H.DB.Exec("update cats set salary=$1 where name=$2", salary, name)
	return err
}

func ListCats() ([]models.Cat, error) {
	H.M.Lock()
	defer H.M.Unlock()
	var cats []models.Cat
	return cats, H.DB.Select(&cats, "select * from cats")
}

func GetCat(name string) (models.Cat, error) {
	exists, err := CatExists(name)
	if err != nil {
		return models.Cat{}, err
	}
	if !exists {
		return models.Cat{}, ErrorCatDoesntExist
	}
	H.M.Lock()
	defer H.M.Unlock()
	var cat models.Cat
	return cat, H.DB.Get(&cat, "select * from cats where name=$1", name)
}

func CatExists(name string) (bool, error) {
	H.M.Lock()
	defer H.M.Unlock()
	var data []models.Cat
	err := H.DB.Select(&data, "select * from cats where name=$1", name)
	if err != nil {
		return false, err
	}
	return len(data) > 0, err
}
