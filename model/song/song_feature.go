package song

import (
	"encoding/json"
	"errors"
	"log/slog"
	"strconv"

	"github.com/ponegraph/backend/exception"
	"github.com/ponegraph/backend/helper"
)

type SongFeature struct {
	SongId           int `json:"songId"`
	BPM              int `json:"bpm"`
	Danceability     int `json:"danceability"`
	Energy           int `json:"energy"`
	Valence          int `json:"valence"`
	Acousticness     int `json:"acousticness"`
	Instrumentalness int `json:"instrumentalness"`
	Liveness         int `json:"liveness"`
	Speechiness      int `json:"speechiness"`
}

type SongFeatureGraphDBResult struct {
	Results struct {
		Bindings []struct {
			SongId struct {
				Value string `json:"value"`
			} `json:"songId"`
			BPM struct {
				Value string `json:"value"`
			} `json:"bpm"`
			Danceability struct {
				Value string `json:"value"`
			} `json:"danceability"`
			Energy struct {
				Value string `json:"value"`
			} `json:"energy"`
			Valence struct {
				Value string `json:"value"`
			} `json:"valence"`
			Acousticness struct {
				Value string `json:"value"`
			} `json:"acousticness"`
			Instrumentalness struct {
				Value string `json:"value"`
			} `json:"instrumentalness"`
			Liveness struct {
				Value string `json:"value"`
			} `json:"liveness"`
			Speechiness struct {
				Value string `json:"value"`
			} `json:"speechiness"`
		} `json:"bindings"`
	} `json:"results"`
}

func ConvertToSongFeatureList(responseBody []byte) ([]SongFeature, error) {
	var result SongFeatureGraphDBResult
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		slog.Error("Failed to execute request", "error", err.Error())
		return nil, errors.New(helper.ErrFailedDatabaseQuery)
	}

	if len(result.Results.Bindings) == 0 {
		return nil, exception.NewNotFoundError(helper.ErrSongNotFound)
	}

	var songList []SongFeature
	for _, binding := range result.Results.Bindings {
		song := SongFeature{}

		song.SongId, _ = strconv.Atoi(binding.SongId.Value)
		song.BPM, _ = strconv.Atoi(binding.BPM.Value)
		song.Danceability, _ = strconv.Atoi(binding.Danceability.Value)
		song.Energy, _ = strconv.Atoi(binding.Energy.Value)
		song.Valence, _ = strconv.Atoi(binding.Valence.Value)
		song.Acousticness, _ = strconv.Atoi(binding.Acousticness.Value)
		song.Instrumentalness, _ = strconv.Atoi(binding.Instrumentalness.Value)
		song.Liveness, _ = strconv.Atoi(binding.Liveness.Value)
		song.Speechiness, _ = strconv.Atoi(binding.Speechiness.Value)

		songList = append(songList, song)
	}
	return songList, nil
}
