package artist

import (
	"encoding/json"
	"errors"

	"github.com/ponegraph/backend/exception"
	"github.com/ponegraph/backend/helper"
)

type ArtistUnit struct {
	Name     string `json:"artistName"`
	ArtistId string `json:"artistId"`
	MbUrl    string `json:"mbUrl"`
}

type ArtistUnitGraphDBResult struct {
	Results struct {
		Bindings []struct {
			Name struct {
				Value string `json:"value"`
			} `json:"artistName"`
			ArtistId struct {
				Value string `json:"value"`
			} `json:"artistId"`
			MbUrl struct {
				Value string `json:"value"`
			} `json:"mbUrl"`
		} `json:"bindings"`
	} `json:"results"`
}

func ConvertToArtistUnitList(responseBody []byte) ([]ArtistUnit, error) {
	var result ArtistUnitGraphDBResult
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, errors.New(helper.ErrFailedDatabaseQuery)
	}

	if len(result.Results.Bindings) == 0 {
		return nil, exception.NewNotFoundError(helper.ErrArtistNotFound)
	}

	var artistUnitList []ArtistUnit
	for _, binding := range result.Results.Bindings {
		artistUnit := ArtistUnit{
			Name:     binding.Name.Value,
			ArtistId: binding.ArtistId.Value,
			MbUrl:    binding.MbUrl.Value,
		}

		artistUnitList = append(artistUnitList, artistUnit)
	}

	return artistUnitList, nil
}
