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

func (c *Client) fetchLocationAreaBatch(url *string) (LocationAreaBatch, error) {

	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return LocationAreaBatch{}, fmt.Errorf("Error creating request: %v", err)
	}


	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaBatch{}, fmt.Errorf("Error request next locations: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaBatch{}, fmt.Errorf("Error reading response body: %v", err)
	}

	var locAreas LocationAreaBatch
	err = json.Unmarshal(data, &locAreas)
	if err != nil {
		return LocationAreaBatch{}, fmt.Errorf("Error decoding response body: %v", err)
	}

	return locAreas, nil
}
