package artist

import (
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/ponegraph/backend/exception"
	"github.com/ponegraph/backend/helper"
)

type ArtistDbpedia struct {
	Description string `json:"description"`
	ExternalUrl string `json:"externalReference"`
	ImageUrl    string `json:"imageUrl"`
}

type ArtistDbpediaResult struct {
	Results struct {
		Bindings []struct {
			Description struct {
				Value string `json:"value"`
			} `json:"description"`
			ExternalUrl struct {
				Value string `json:"value"`
			} `json:"externalReference"`
			ImageUrl struct {
				Value string `json:"value"`
			} `json:"imageUrl"`
		} `json:"bindings"`
	} `json:"results"`
}

func ConvertToArtistDbpedia(responseBody []byte) (*ArtistDbpedia, error) {
	var result ArtistDbpediaResult
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, errors.New(helper.ErrFailedDatabaseQuery)
	}

	if len(result.Results.Bindings) == 0 {
		return nil, exception.NewNotFoundError("artist not found in remote database")
	}

	binding := result.Results.Bindings[0]

	artistDbpedia := ArtistDbpedia{
		Description: binding.Description.Value,
		ExternalUrl: binding.ExternalUrl.Value,
		ImageUrl:    binding.ImageUrl.Value,
	}

	return &artistDbpedia, nil
}

type ArtistImageResult struct {
	Results struct {
		Bindings []struct {
			ImageUrl struct {
				Value string `json:"value"`
			} `json:"imageUrl"`
		} `json:"bindings"`
	} `json:"results"`
}

func ConvertToArtistImage(responseBody []byte) (string, error) {
	var result ArtistImageResult
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		slog.Error("Failed to execute request", "error", err.Error())
		return "", errors.New(helper.ErrFailedDatabaseQuery)
	}

	if len(result.Results.Bindings) == 0 {
		return "", nil
	}

	binding := result.Results.Bindings[0]

	return binding.ImageUrl.Value, nil
}
