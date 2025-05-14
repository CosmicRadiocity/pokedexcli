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

type LocationAreaDetails struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) FetchLocationAreaBatch(url string) (LocationAreaBatch, error) {

	var data []byte
	var err error

	entry, ok := c.cache.Get(url)
	if ok {
		data = entry
	} else {
		if url == "" {
			url = baseURL + "/location-area"
		}

		data, err = c.FetchDataFromUrl(url)
		if err != nil {
			return LocationAreaBatch{}, err
		}
	}

	var locAreas LocationAreaBatch
	err = json.Unmarshal(data, &locAreas)
	if err != nil {
		return LocationAreaBatch{}, fmt.Errorf("Error decoding response body: %v", err)
	}

	return locAreas, nil
}

func (c *Client) FetchLocationAreaDetails(name string) (LocationAreaDetails, error) {
	var data []byte
	var err error

	url := baseURL + "/location-area/" + name

	entry, ok := c.cache.Get(url)
	if ok {
		data = entry
	} else {
		data, err = c.FetchDataFromUrl(url)
		if err != nil {
			return LocationAreaDetails{}, err
		}
	}

	var details LocationAreaDetails
	err = json.Unmarshal(data, &details)
	if err != nil {
		return LocationAreaDetails{}, fmt.Errorf("Error decoding response body: %v", err)
	}

	return details, nil
}

func (c *Client) FetchDataFromUrl(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Error creating request: %v", err)
	}


	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("Error request next locations: %v", err)
	}
	defer res.Body.Close()

	var data []byte

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Error reading response body: %v", err)
	}

	c.cache.Add(url, data)
	return data, nil
}
