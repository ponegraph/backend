package web

import (
	artistModel "github.com/ponegraph/backend/model/artist"
)

type ArtistListResponse struct {
	Artists []artistModel.ArtistUnit `json:"artists"`
}
