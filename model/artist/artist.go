package artist

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/ponegraph/backend/exception"
	"github.com/ponegraph/backend/helper"
)

type Artist struct {
	Name                 string   `json:"artistName"`
	ArtistId             string   `json:"artistId"`
	MbUrl                string   `json:"mbUrl"`
	Country              string   `json:"countryName"`
	TotalLastfmListeners int      `json:"totalLastfmListeners"`
	TotalLastfmScrobbles int      `json:"totalLastfmScrobbles"`
	Tags                 []string `json:"tags"`

	// from DBpedia
	AdditioanlInfo ArtistDbpedia `json:"additionalInfo"`
}

type ArtistGraphDBResult struct {
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
			Country struct {
				Value string `json:"value"`
			} `json:"countryName"`
			TotalLastfmListeners struct {
				Value string `json:"value"`
			} `json:"totalLastfmListeners"`
			TotalLastfmScrobbles struct {
				Value string `json:"value"`
			} `json:"totalLastfmScrobbles"`
			Tags struct {
				Value string `json:"value"`
			} `json:"tags"`
		} `json:"bindings"`
	} `json:"results"`
}

func ConvertToArtist(responseBody []byte) (*Artist, error) {
	var result ArtistGraphDBResult
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, errors.New(helper.ErrFailedDatabaseQuery)
	}

	if len(result.Results.Bindings) == 0 {
		return nil, exception.NewNotFoundError(helper.ErrArtistNotFound)
	}

	binding := result.Results.Bindings[0]

	artist := Artist{
		Name:     binding.Name.Value,
		ArtistId: binding.ArtistId.Value,
		MbUrl:    binding.MbUrl.Value,
		Country:  binding.Country.Value,
	}

	artist.TotalLastfmListeners, _ = strconv.Atoi(binding.TotalLastfmListeners.Value)
	artist.TotalLastfmScrobbles, _ = strconv.Atoi(binding.TotalLastfmScrobbles.Value)

	if binding.Tags.Value != "" {
		artist.Tags = strings.Split(binding.Tags.Value, ", ")
	}

	return &artist, nil
}
