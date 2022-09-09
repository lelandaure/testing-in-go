package util

import (
	"encoding/json"
	"github.com/lelandaure/testing-in-go/models"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestParsePokemonSuccess(t *testing.T) {
	assert := require.New(t)

	expected := apiResponse(assert)

	pokeapiResponse := pokeApiResponse(assert)

	actual, err := ParsePokemon(pokeapiResponse)
	assert.NoError(err)

	assert.Equal(expected, actual, "This should be equal")
}

func TestParsePokemonTypeNotFound(t *testing.T) {
	assert := require.New(t)

	response := pokeApiResponse(assert)

	response.PokemonType = []models.PokemonType{}

	_, err := ParsePokemon(response)
	//assert.NotNil(err)
	assert.EqualError(ErrNotFoundPokemonType, err.Error())
}

func TestParsePokemonTypeNotFoundName(t *testing.T) {
	assert := require.New(t)

	pokeapiResponse := pokeApiResponse(assert)
	pokeapiResponse.PokemonType[0].RefType.Name = ""

	_, err := ParsePokemon(pokeapiResponse)

	assert.EqualError(ErrNotFoundPokemonTypeName, err.Error())
}

func apiResponse(assert *require.Assertions) models.Pokemon {
	apiResponseJson, err := os.ReadFile("sample/api_response.json")
	assert.NoError(err)

	var pokemon models.Pokemon

	err = json.Unmarshal(apiResponseJson, &pokemon)
	assert.NoError(err)

	return pokemon
}

func pokeApiResponse(assert *require.Assertions) models.PokeApiPokemonResponse {
	pokeApiResponseJson, err := os.ReadFile("sample/pokeapi_response.json")
	assert.NoError(err)

	var pokeApiPokemonResponse models.PokeApiPokemonResponse

	err = json.Unmarshal(pokeApiResponseJson, &pokeApiPokemonResponse)
	assert.NoError(err)

	return pokeApiPokemonResponse
}
