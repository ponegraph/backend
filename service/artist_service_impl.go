package service

import (
	artistModel "github.com/ponegraph/backend/model/artist"
	"log/slog"

	"github.com/ponegraph/backend/helper"
	"github.com/ponegraph/backend/model/web"
	"github.com/ponegraph/backend/repository"
)

type ArtistServiceImpl struct {
	ArtistRepository repository.ArtistRepository
	SongRepository   repository.SongRepository
	logger           *slog.Logger
}

func NewArtistService(
	songRepository repository.SongRepository, artistRepository repository.ArtistRepository, logger *slog.Logger,
) ArtistService {
	return &ArtistServiceImpl{
		SongRepository:   songRepository,
		ArtistRepository: artistRepository,
		logger:           logger,
	}
}

func (service *ArtistServiceImpl) GetArtistDetail(artistId string) web.ArtistDetailResponse {
	artist, err := service.ArtistRepository.GetByArtistId(artistId)
	helper.PanicIfError(err)

	artistSongs, err := service.SongRepository.GetAllUnitByArtistId(artistId)
	helper.PanicIfError(err)

	artistInfoFromDbpedia, _ := service.ArtistRepository.GetInfoFromDbpedia(artist.MbUrl)

	if artistInfoFromDbpedia == nil {
		artist.AdditioanlInfo = artistModel.ArtistDbpedia{}
	} else {
		artist.AdditioanlInfo = *artistInfoFromDbpedia
	}
	response := web.ArtistDetailResponse{
		Artist: *artist,
		Songs:  artistSongs,
	}
	return response
}

func (service *ArtistServiceImpl) SearchArtistByTag(tag string) web.ArtistListResponse {
	artistUnitList, err := service.ArtistRepository.GetAllUnitByTag(tag)
	helper.PanicIfError(err)

	response := web.ArtistListResponse{
		Artists: artistUnitList,
	}
	return response
}

func (service *ArtistServiceImpl) SearchArtistByName(artistName string) web.ArtistListResponse {
	artistUnitList, err := service.ArtistRepository.GetAllUnitByName(artistName)
	helper.PanicIfError(err)

	response := web.ArtistListResponse{
		Artists: artistUnitList,
	}
	return response
}

func (service *ArtistServiceImpl) GetTopRank() web.ArtistListResponse {
	artistUnitList, err := service.ArtistRepository.GetTopRank()
	helper.PanicIfError(err)

	response := web.ArtistListResponse{
		Artists: artistUnitList,
	}
	return response
}
