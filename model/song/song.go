package song

import (
	"encoding/json"
	"errors"
	"github.com/ponegraph/backend/exception"
	"github.com/ponegraph/backend/helper"
	"strconv"
)

type Song struct {
	Name        string `json:"songName"`
	SongId      int    `json:"songId"`
	ReleaseDate string `json:"releaseDate"`

	BPM           int    `json:"bpm"`
	Key           string `json:"key"`
	Mode          string `json:"mode"`
	SpotifyStream int    `json:"spotifyStream"`

	SpotifyPlaylistCount int `json:"spotifyPlaylistCount"`
	ApplePlaylistCount   int `json:"applePlaylistCount"`
	DeezerPlaylistCount  int `json:"deezerPlaylistCount"`

	SpotifyChart int `json:"spotifyChart"`
	AppleChart   int `json:"appleChart"`
	DeezerChart  int `json:"deezerChart"`
	ShazamChart  int `json:"shazamChart"`

	Danceability     int `json:"danceability"`
	Energy           int `json:"energy"`
	Valence          int `json:"valence"`
	Acousticness     int `json:"acousticness"`
	Instrumentalness int `json:"instrumentalness"`
	Liveness         int `json:"liveness"`
	Speechiness      int `json:"speechiness"`
}

type SongGraphDBResult struct {
	Results struct {
		Bindings []struct {
			Name struct {
				Value string `json:"value"`
			} `json:"songName"`
			ReleaseDate struct {
				Value string `json:"value"`
			} `json:"releaseDate"`
			BPM struct {
				Value string `json:"value"`
			} `json:"bpm"`
			Key struct {
				Value string `json:"value"`
			} `json:"key"`
			Mode struct {
				Value string `json:"value"`
			} `json:"mode"`
			SpotifyStream struct {
				Value string `json:"value"`
			} `json:"spotifyStream"`
			SpotifyPlaylistCount struct {
				Value string `json:"value"`
			} `json:"spotifyPlaylistCount"`
			ApplePlaylistCount struct {
				Value string `json:"value"`
			} `json:"applePlaylistCount"`
			DeezerPlaylistCount struct {
				Value string `json:"value"`
			} `json:"deezerPlaylistCount"`
			SpotifyChart struct {
				Value string `json:"value"`
			} `json:"spotifyChart"`
			AppleChart struct {
				Value string `json:"value"`
			} `json:"appleChart"`
			DeezerChart struct {
				Value string `json:"value"`
			} `json:"deezerChart"`
			ShazamChart struct {
				Value string `json:"value"`
			} `json:"shazamChart"`
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

func ConvertToSong(responseBody []byte) (*Song, error) {
	var result SongGraphDBResult
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, errors.New(helper.ErrFailedDatabaseQuery)
	}

	if len(result.Results.Bindings) == 0 {
		return nil, exception.NewNotFoundError(helper.ErrSongNotFound)
	}

	binding := result.Results.Bindings[0]

	song := Song{
		Name:        binding.Name.Value,
		ReleaseDate: binding.ReleaseDate.Value,
		Key:         binding.Key.Value,
	}

	song.BPM, _ = strconv.Atoi(binding.BPM.Value)
	song.SpotifyStream, _ = strconv.Atoi(binding.SpotifyStream.Value)
	song.SpotifyPlaylistCount, _ = strconv.Atoi(binding.SpotifyPlaylistCount.Value)
	song.ApplePlaylistCount, _ = strconv.Atoi(binding.ApplePlaylistCount.Value)
	song.DeezerPlaylistCount, _ = strconv.Atoi(binding.DeezerPlaylistCount.Value)
	song.SpotifyChart, _ = strconv.Atoi(binding.SpotifyChart.Value)
	song.AppleChart, _ = strconv.Atoi(binding.AppleChart.Value)
	song.DeezerChart, _ = strconv.Atoi(binding.DeezerChart.Value)
	song.ShazamChart, _ = strconv.Atoi(binding.ShazamChart.Value)
	song.Danceability, _ = strconv.Atoi(binding.Danceability.Value)
	song.Energy, _ = strconv.Atoi(binding.Energy.Value)
	song.Valence, _ = strconv.Atoi(binding.Valence.Value)
	song.Acousticness, _ = strconv.Atoi(binding.Acousticness.Value)
	song.Instrumentalness, _ = strconv.Atoi(binding.Instrumentalness.Value)
	song.Liveness, _ = strconv.Atoi(binding.Liveness.Value)
	song.Speechiness, _ = strconv.Atoi(binding.Speechiness.Value)

	return &song, nil
}
