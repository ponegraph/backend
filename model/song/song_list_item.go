package song

import (
	artistModel "github.com/ponegraph/backend/model/artist"
)

type SongListItem struct {
	Name        string                   `json:"songName"`
	SongId      int                      `json:"songId"`
	ReleaseDate string                   `json:"releaseDate"`
	Artists     []artistModel.ArtistUnit `json:"artists"`
}

func NewSongListItem(songUnit SongUnit, artist []artistModel.ArtistUnit) *SongListItem {
	return &SongListItem{
		Name:        songUnit.Name,
		SongId:      songUnit.SongId,
		ReleaseDate: songUnit.ReleaseDate,
		Artists:     artist,
	}
}
