package service

import (
	"github.com/ponegraph/backend/model/web"
)

type SongService interface {
	GetSongDetail(songId int) web.SongDetailResponse
	GetTopRank() web.SongListResponse
	GetSongByArtistId(artistId string) web.SongArtistResponse
}
