package repository

import (
	"log/slog"

	"github.com/ponegraph/backend/helper"
	songModel "github.com/ponegraph/backend/model/song"
)

type SongRepositoryImpl struct {
	logger *slog.Logger
}

func NewSongRepository(logger *slog.Logger) SongRepository {
	return &SongRepositoryImpl{
		logger: logger,
	}
}

func (repository *SongRepositoryImpl) GetById(songId int) (*songModel.Song, error) {
	query := helper.GetSongByIdQuery(songId)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	song, err := songModel.ConvertToSong(responseBody)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (repository *SongRepositoryImpl) GetTopRank() ([]songModel.SongUnit, error) {
	query := helper.GetTopKSongUnitQuery(10)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	songList, err := songModel.ConvertToSongUnitList(responseBody)
	if err != nil {
		return nil, err
	}

	return songList, nil
}

func (repository *SongRepositoryImpl) GetAllSongFeature() ([]songModel.SongFeature, error) {
	query := helper.GetAllSongFeatureQuery()
	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	songList, err := songModel.ConvertToSongFeatureList(responseBody)
	if err != nil {
		return nil, err
	}

	return songList, nil
}

func (repository *SongRepositoryImpl) GetAllSongIdFromSameArtist(songId int) ([]int, error) {
	query := helper.GetAllSongIdFromSameArtistQuery(songId)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	songIdList, err := songModel.ConvertToSongIdList(responseBody)
	if err != nil {
		return nil, err
	}

	return songIdList, nil
}

func (repository *SongRepositoryImpl) GetUnitBySongId(songId int) (*songModel.SongUnit, error) {
	query := helper.GetSongUnitByIdQuery(songId)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	song, err := songModel.ConvertToSongUnit(responseBody)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (repository *SongRepositoryImpl) GetAllUnitByArtistId(artistId string) ([]songModel.SongUnit, error) {
	query := helper.GetAllSongUnitFromArtistIdQuery(artistId)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	songList, err := songModel.ConvertToSongUnitList(responseBody)
	if err != nil {
		return nil, err
	}

	return songList, nil
}

func (repository *SongRepositoryImpl) GetAllUnitByName(name string) ([]songModel.SongUnit, error) {
	query := helper.GetAllSongUnitByNameQuery(name)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	songList, err := songModel.ConvertToSongUnitList(responseBody)
	if err != nil {
		return nil, err
	}

	return songList, nil
}
