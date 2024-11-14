package repository

import (
	artistModel "github.com/ponegraph/backend/model/artist"
)

type ArtistRepository interface {
	GetByArtistId(artistId string) (*artistModel.Artist, error)
	GetInfoFromDbpedia(artistId string) (*artistModel.ArtistDbpedia, error)
	GetAllUnitBySongId(songId int) ([]artistModel.ArtistUnit, error)
	GetAllUnitByTag(tag string) ([]artistModel.ArtistUnit, error)
	GetAllUnitByName(name string) ([]artistModel.ArtistUnit, error)
}
