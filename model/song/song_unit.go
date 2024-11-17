package song

import (
	"encoding/json"
	"errors"
	"log/slog"
	"strconv"

	"github.com/ponegraph/backend/exception"
	"github.com/ponegraph/backend/helper"
)

type SongUnit struct {
	Name        string `json:"songName"`
	SongId      int    `json:"songId"`
	ReleaseDate string `json:"releaseDate"`
}

type SongUnitGraphDBResult struct {
	Results struct {
		Bindings []struct {
			Name struct {
				Value string `json:"value"`
			} `json:"songName"`
			SongId struct {
				Value string `json:"value"`
			} `json:"songId"`
			ReleaseDate struct {
				Value string `json:"value"`
			} `json:"releaseDate"`
		} `json:"bindings"`
	} `json:"results"`
}

func ConvertToSongUnit(responseBody []byte) (*SongUnit, error) {
	var result SongUnitGraphDBResult
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		slog.Error("Failed to execute request", "error", err.Error())
		return nil, errors.New(helper.ErrFailedDatabaseQuery)
	}

	if len(result.Results.Bindings) == 0 {
		return nil, exception.NewNotFoundError(helper.ErrSongNotFound)
	}

	binding := result.Results.Bindings[0]
	songUnit := SongUnit{
		Name:        binding.Name.Value,
		ReleaseDate: binding.ReleaseDate.Value,
	}
	songUnit.SongId, _ = strconv.Atoi(binding.SongId.Value)

	return &songUnit, nil
}

func ConvertToSongUnitList(responseBody []byte) ([]SongUnit, error) {
	var result SongUnitGraphDBResult
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		slog.Error("Failed to execute request", "error", err.Error())
		return nil, errors.New(helper.ErrFailedDatabaseQuery)
	}

	if len(result.Results.Bindings) == 0 {
		return nil, exception.NewNotFoundError(helper.ErrSongNotFound)
	}

	var songUnitList []SongUnit
	for _, binding := range result.Results.Bindings {
		songUnit := SongUnit{
			Name:        binding.Name.Value,
			ReleaseDate: binding.ReleaseDate.Value,
		}
		songUnit.SongId, _ = strconv.Atoi(binding.SongId.Value)
		songUnitList = append(songUnitList, songUnit)
	}

	return songUnitList, nil
}

type SongIdGraphDBResult struct {
	Results struct {
		Bindings []struct {
			SongId struct {
				Value string `json:"value"`
			} `json:"songId"`
		} `json:"bindings"`
	} `json:"results"`
}

func ConvertToSongIdList(responseBody []byte) ([]int, error) {
	var result SongIdGraphDBResult
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		slog.Error("Failed to execute request", "error", err.Error())
		return nil, errors.New(helper.ErrFailedDatabaseQuery)
	}

	if len(result.Results.Bindings) == 0 {
		return nil, exception.NewNotFoundError(helper.ErrSongNotFound)
	}

	var songIdList []int
	for _, binding := range result.Results.Bindings {
		songId, _ := strconv.Atoi(binding.SongId.Value)

		songIdList = append(songIdList, songId)
	}

	return songIdList, nil
}
