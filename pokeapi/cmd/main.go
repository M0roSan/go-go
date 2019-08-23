package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
    "encoding/json"
	
)

// Response struct to map the entire response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// Pokemon struc to map every pokemon to
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// PokemonSpecies a struct to map our Pokemon species which includes its name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	resp, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
    json.Unmarshal(respData, &responseObject)

    fmt.Println(responseObject.Name)
    fmt.Println(len(responseObject.Pokemon))

    for i := 0; i < len(responseObject.Pokemon); i++ {
        fmt.Println(responseObject.Pokemon[i].Species.Name)
    }

}