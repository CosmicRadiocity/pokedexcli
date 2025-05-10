package pokeapi

import ( 
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

type LocationAreaBatch struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func fetchLocationAreaBatch(url string, config *config) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return []string{}, fmt.Errorf("Error request next locations: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, fmt.Errorf("Error reading response body: %v", err)
	}

	var locAreas LocationAreaBatch
	err = json.Unmarshal(data, &locAreas)
	if err != nil {
		return []string{}, fmt.Errorf("Error decoding response body: %v", err)
	}
	config.Next = locAreas.Next
	config.Previous = locAreas.Next
	results := []string{}
	for _, locArea := range locAreas.Results {
		results = append(results, locArea.Name)
	}
	return results, nil
}
