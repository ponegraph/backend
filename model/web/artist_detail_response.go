package web

import (
	artistModel "github.com/ponegraph/backend/model/artist"
	songModel "github.com/ponegraph/backend/model/song"
)

type ArtistDetailResponse struct {
	Artist artistModel.Artist   `json:"artist"`
	Songs  []songModel.SongUnit `json:"songs"`
}
