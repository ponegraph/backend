package web

import (
	songModel "github.com/ponegraph/backend/model/song"
)

type SongArtistResponse struct {
	Songs []songModel.SongUnit `json:"songs"`
}
