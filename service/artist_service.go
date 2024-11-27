package service

import (
	"github.com/ponegraph/backend/model/web"
)

type ArtistService interface {
	GetArtistDetail(artistId string) web.ArtistDetailResponse
	SearchArtistByTag(tag string) web.ArtistListResponse
	SearchArtistByName(tag string) web.ArtistListResponse
	GetTopRank() web.ArtistListResponse
}
