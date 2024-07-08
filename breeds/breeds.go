package breeds

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

// In this package we will cache breeds so it would be faster to check for correct breed

type MBreeds struct {
	M   sync.RWMutex
	Map map[string]bool
}

func (m *MBreeds) CheckExists(key string) bool {
	m.M.RLock()
	defer m.M.RUnlock()
	_, ok := m.Map[key]
	return ok
}

func (m *MBreeds) GetAll() []string {
	m.M.RLock()
	defer m.M.RUnlock()
	var res []string
	for k := range m.Map {
		res = append(res, k)
	}
	return res
}

var Breeds MBreeds

type BreedData struct {
	Name string `json:"name"`
}

func init() {
	Breeds.M.Lock()
	Breeds.Map = make(map[string]bool)
	Breeds.M.Unlock()
	log.Println("Loading breeds")
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds") //hardcoded because it's not gonna change
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("The status code is %d, expected 200\n", resp.StatusCode)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var breeds []BreedData
	err = json.Unmarshal(b, &breeds)
	if err != nil {
		log.Fatalln(err)
	}
	Breeds.M.Lock()
	defer Breeds.M.Unlock()
	for _, breed := range breeds {
		Breeds.Map[strings.ToLower(breed.Name)] = true
	}
}

func ValidateBreed(breed string) bool {
	return Breeds.CheckExists(strings.ToLower(breed))
}
