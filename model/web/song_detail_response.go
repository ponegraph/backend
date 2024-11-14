package web

import (
	artistModel "github.com/ponegraph/backend/model/artist"
	songModel "github.com/ponegraph/backend/model/song"
)

type SongDetailResponse struct {
	Song             songModel.Song           `json:"song"`
	Artists          []artistModel.ArtistUnit `json:"artists"`
	RecommendedSongs []songModel.SongListItem `json:"recommendedSongs"`
}
