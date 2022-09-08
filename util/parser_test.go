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

	body, err := os.ReadFile("sample/pokeapi_response.json")
	assert.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &response)
	assert.NoError(err)

	parsedPokemon, err := ParsePokemon(response)
	assert.NoError(err)

	body, err = os.ReadFile("sample/api_response.json")
	assert.NoError(err)

	var expectedPokemon models.Pokemon

	err = json.Unmarshal(body, &expectedPokemon)
	assert.NoError(err)

	assert.Equal(expectedPokemon, parsedPokemon, "This should be equal")
}
