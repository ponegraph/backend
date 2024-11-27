package repository

import (
	"log/slog"

	"github.com/ponegraph/backend/helper"
	artistModel "github.com/ponegraph/backend/model/artist"
)

type ArtistRepositoryImpl struct {
	logger *slog.Logger
}

func NewArtistRepository(logger *slog.Logger) ArtistRepository {
	return &ArtistRepositoryImpl{
		logger: logger,
	}
}

func (repository *ArtistRepositoryImpl) GetAllUnitBySongId(songId int) ([]artistModel.ArtistUnit, error) {
	query := helper.GetAllArtistUnitBySongIdQuery(songId)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	artistUnitList, err := artistModel.ConvertToArtistUnitList(responseBody)
	if err != nil {
		return nil, err
	}

	return artistUnitList, nil
}

func (repository *ArtistRepositoryImpl) GetByArtistId(artistId string) (*artistModel.Artist, error) {
	query := helper.GetArtistByIdQuery(artistId)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	artist, err := artistModel.ConvertToArtist(responseBody)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (repository *ArtistRepositoryImpl) GetAllUnitByTag(tag string) ([]artistModel.ArtistUnit, error) {
	query := helper.GetAllArtistUnitByTagQuery(tag)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	artistUnitList, err := artistModel.ConvertToArtistUnitList(responseBody)
	if err != nil {
		return nil, err
	}

	return artistUnitList, nil
}

func (repository *ArtistRepositoryImpl) GetAllUnitByName(name string) ([]artistModel.ArtistUnit, error) {
	query := helper.GetAllArtistUnitByNameQuery(name)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	artistUnitList, err := artistModel.ConvertToArtistUnitList(responseBody)
	if err != nil {
		return nil, err
	}

	return artistUnitList, nil
}

func (repository *ArtistRepositoryImpl) GetInfoFromDbpedia(mbUrl string) (*artistModel.ArtistDbpedia, error) {
	query := helper.GetArtistInfoFromDbpediaQuery(mbUrl)

	responseBody, err := helper.ExecuteDBpediaQuery(query)
	if err != nil {
		return nil, err
	}

	artistDbpedia, err := artistModel.ConvertToArtistDbpedia(responseBody)
	if err != nil {
		return nil, err
	}

	return artistDbpedia, nil
}

func (repository *ArtistRepositoryImpl) GetTopRank() ([]artistModel.ArtistUnit, error) {
	query := helper.GetTopKArtistUnitQuery(10)

	responseBody, err := helper.ExecuteGraphDBQuery(query)
	if err != nil {
		return nil, err
	}

	artistUnitList, err := artistModel.ConvertToArtistUnitList(responseBody)
	if err != nil {
		return nil, err
	}

	return artistUnitList, nil
}
