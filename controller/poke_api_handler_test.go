package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/lelandaure/testing-in-go/models"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"testing"
)

func TestGetPokemonFromPokeApiFirstWay(t *testing.T) {
	assert := require.New(t)
	id := "bulbasaur"
	tt := []struct {
		testName         string
		idOrName         string
		expectedResponse models.PokeApiPokemonResponse
		expectedError    error
	}{
		{"Success", id, expectedResponseOk(assert), nil},
		{"httpGetError", id, models.PokeApiPokemonResponse{}, errors.New("httpGetError")},
		{"StatusNotFound", id, models.PokeApiPokemonResponse{}, ErrPokemonNotFound},
		{"InternalServerError", id, models.PokeApiPokemonResponse{}, ErrPokeApiFailure},
	}

	for _, test := range tt {
		t.Run(test.testName, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				http.MethodGet,
				fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id),
				func(request *http.Request) (*http.Response, error) {
					responder, err := httpmock.NewJsonResponse(http.StatusOK, test.expectedResponse)
					if test.expectedError != nil && test.expectedError.Error() == "httpGetError" {
						return httpmock.NewStringResponse(http.StatusInternalServerError, ""), test.expectedError
					}
					if test.expectedError != nil && errors.Is(test.expectedError, ErrPokemonNotFound) {
						return httpmock.NewStringResponse(http.StatusNotFound, ""), nil
					}
					if err != nil || test.expectedError != nil {
						return httpmock.NewStringResponse(http.StatusInternalServerError, ""), nil
					}
					return responder, nil
				},
			)

			actual, err := GetPokemonFromPokeApi(test.idOrName)
			assert.ErrorIs(err, test.expectedError)
			assert.Equal(test.expectedResponse, actual)

		})
	}
}

func TestGetPokemonFromPokeApiSecondWay(t *testing.T) {
	assert := require.New(t)
	for scenario, fn := range map[string]func(assert *require.Assertions){
		"Success":      testSuccess,
		"httpGetError": testHttpGetError,
		//"StatusNotFound": testStatusNotFound,
		//"InternalServcerError": testInternalServcerError,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(assert)
		})
	}
}

func testHttpGetError(assert *require.Assertions) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", "bulbasaur"),
		httpmock.NewErrorResponder(errors.New("dummy error")),
	)

	_, err := GetPokemonFromPokeApi("bulbasaur")

	assert.Error(err)
}

func testSuccess(assert *require.Assertions) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	expectedResponse := expectedResponseOk(assert)
	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", "bulbasaur"),
		httpmock.NewJsonResponderOrPanic(http.StatusOK, expectedResponse),
	)

	actual, err := GetPokemonFromPokeApi("bulbasaur")
	assert.NoError(err)
	assert.Equal(expectedResponse, actual)
}

func expectedResponseOk(assert *require.Assertions) models.PokeApiPokemonResponse {
	bodyPokeApiResponse, err := os.ReadFile("sample/poke_api_response.json")
	assert.NoError(err)

	var pokeApiResponseOk models.PokeApiPokemonResponse

	err = json.Unmarshal(bodyPokeApiResponse, &pokeApiResponseOk)
	assert.NoError(err)

	return pokeApiResponseOk
}
