package breeds

import "testing"

func TestMBreeds_GetAll(t *testing.T) {
	res := Breeds.GetAll()
	if res == nil {
		t.Error("Empty res")
		t.FailNow()
	}
	t.Log(res)
}

func TestValidateBreed(t *testing.T) {
	breed := "malayan"
	breed2 := "12345"
	if !ValidateBreed(breed) {
		t.Error("Expected breed doesn't exist")
		t.FailNow()
	}
	if ValidateBreed(breed2) {
		t.Error("Unexpected breed exists")
		t.FailNow()
	}
}
