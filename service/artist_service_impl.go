package service

import (
	"github.com/ponegraph/backend/helper"
	"github.com/ponegraph/backend/model/web"
	"github.com/ponegraph/backend/repository"
)

type ArtistServiceImpl struct {
	ArtistRepository repository.ArtistRepository
	SongRepository   repository.SongRepository
}

func NewArtistService(songRepository repository.SongRepository, artistRepository repository.ArtistRepository) ArtistService {
	return &ArtistServiceImpl{
		SongRepository:   songRepository,
		ArtistRepository: artistRepository,
	}
}

func (service *ArtistServiceImpl) GetArtistDetail(artistId string) web.ArtistDetailResponse {
	artist, err := service.ArtistRepository.GetByArtistId(artistId)
	helper.PanicIfError(err)

	artistSongs, err := service.SongRepository.GetAllUnitByArtistId(artistId)
	helper.PanicIfError(err)

	artistInfoFromDbpedia, _ := service.ArtistRepository.GetInfoFromDbpedia(artist.MbUrl)

	artist.AdditioanlInfo = *artistInfoFromDbpedia
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
