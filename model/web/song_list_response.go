package web

import (
	songModel "github.com/ponegraph/backend/model/song"
)

type SongListResponse struct {
	Songs []songModel.SongListItem `json:"songs"`
}
