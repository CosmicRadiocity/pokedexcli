package pokeapi

import(
	"sync"
	"fmt"
)

type Pokedex struct{
	entries map[string]Pokemon
	mu sync.Mutex
}

func NewPokedex() Pokedex {
	return Pokedex{
		entries: make(map[string]Pokemon),
	}
}

func (p *Pokedex) AddPokemon(name string, pokemon Pokemon) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.entries[name]; !ok {
		fmt.Printf("%s was added to your Pokedex!\n", name)
		p.entries[name] = pokemon
	}
}

func (p *Pokedex) GetPokemon(name string) (Pokemon, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	pokemon, ok := p.entries[name]
	return pokemon, ok
}