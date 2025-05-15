package pokeapi

import (
	"net/http"
	"time"

	"github.com/CosmicRadiocity/pokedexcli/internal/pokecache"
)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
	pokedex Pokedex
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokedex: NewPokedex(),
	}
}