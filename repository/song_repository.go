package repository

import (
	songModel "github.com/ponegraph/backend/model/song"
)

type SongRepository interface {
	GetUnitBySongId(songId int) (*songModel.SongUnit, error)
	GetById(songId int) (*songModel.Song, error)
	GetTopRank() ([]songModel.SongUnit, error)
	GetAllSongFeature() ([]songModel.SongFeature, error)
	GetAllSongIdFromSameArtist(songId int) ([]int, error)
	GetAllUnitByArtistId(artistId string) ([]songModel.SongUnit, error)
}
