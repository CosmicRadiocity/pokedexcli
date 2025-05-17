package pokeapi

import ( 
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

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
		return LocationAreaDetails{}, fmt.Errorf("%s is not a valid area name", name)
	}

	return details, nil
}

func (c *Client) FetchPokemon(name string) (Pokemon, error) {
	var data []byte
	var err error

	url := baseURL + "/pokemon/" + name

	entry, ok := c.cache.Get(url)
	if ok {
		data = entry
	} else {
		data, err = c.FetchDataFromUrl(url)
		if err != nil {
			return Pokemon{}, err
		}
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("%s is not a valid pokemon name", name)
	}

	return pokemon, nil
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

// Pokedex access functions

func (c *Client) AddPokemonToPokedex(name string, pokemon Pokemon) {
	c.pokedex.AddPokemon(name, pokemon)
}

func (c *Client) GetPokemonFromPokedex(name string) (Pokemon, bool) {
	return c.pokedex.GetPokemon(name)
}

func (c *Client) GetAllPokemonFromPokedex() ([]Pokemon) {
	return c.pokedex.GetAllPokemon()
}
