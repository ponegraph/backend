package service

import (
	"errors"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/ponegraph/backend/helper"
	songModel "github.com/ponegraph/backend/model/song"
	"github.com/ponegraph/backend/model/web"
	"github.com/ponegraph/backend/repository"
	"math"
	"slices"
	"sort"
	"strconv"
)

type SongServiceImpl struct {
	SongRepository     repository.SongRepository
	ArtistRepository   repository.ArtistRepository
	SongSimilarityData dataframe.DataFrame
	SongRecommendation map[int][]int
}

func NewSongService(songRepository repository.SongRepository, artistRepository repository.ArtistRepository) SongService {
	newSongService := &SongServiceImpl{
		SongRepository:     songRepository,
		ArtistRepository:   artistRepository,
		SongRecommendation: make(map[int][]int),
	}

	newSongService.InitSongSimilarityData()
	return newSongService

}

func (service *SongServiceImpl) InitSongSimilarityData() {
	songList, err := service.SongRepository.GetAllSongFeature()
	helper.PanicIfError(err)

	df := dataframe.LoadStructs(songList)
	cols := df.Names()

	// Drop SongId column
	columnsToDrop := map[string]bool{"SongId": true}
	var featureCols []string
	for _, col := range cols {
		if !columnsToDrop[col] {
			featureCols = append(featureCols, col)
		}
	}

	dfFeature := df.Select(featureCols)

	songFeatureMatrix := helper.DataframeToMatrix(dfFeature)
	similarityMatrix := helper.CalculateCosSimilarity(songFeatureMatrix)

	simDF := dataframe.LoadMatrix(similarityMatrix)
	songIds := df.Col("SongId").Records()

	err = simDF.SetNames(songIds...)
	if err != nil {
		panic(errors.New("failed to initialize song similarity data"))
	}

	simDF = simDF.Mutate(series.New(songIds, series.Int, "SongId"))
	service.SongSimilarityData = simDF
}

func (service *SongServiceImpl) GetSongDetail(songId int) web.SongDetailResponse {
	song, err := service.SongRepository.GetById(songId)
	helper.PanicIfError(err)

	songArtist, err := service.ArtistRepository.GetAllUnitBySongId(songId)
	helper.PanicIfError(err)

	recommendedSongId := service.getTopSimilarSong(songId)

	var recommendedSong []songModel.SongListItem
	for _, songId := range recommendedSongId {
		recommendedSongUnit, err := service.SongRepository.GetUnitBySongId(songId)
		helper.PanicIfError(err)

		recommendedSongArtist, err := service.ArtistRepository.GetAllUnitBySongId(songId)
		helper.PanicIfError(err)

		songListItem := songModel.NewSongListItem(*recommendedSongUnit, recommendedSongArtist)

		recommendedSong = append(recommendedSong, *songListItem)
	}

	response := web.SongDetailResponse{
		Song:             *song,
		Artists:          songArtist,
		RecommendedSongs: recommendedSong,
	}

	return response
}

func (service *SongServiceImpl) GetTopRank() web.SongListResponse {
	songUnitList, err := service.SongRepository.GetTopRank()
	helper.PanicIfError(err)

	var songList []songModel.SongListItem
	for _, songUnit := range songUnitList {
		artist, err := service.ArtistRepository.GetAllUnitBySongId(songUnit.SongId)
		helper.PanicIfError(err)

		songListItem := songModel.NewSongListItem(songUnit, artist)
		songList = append(songList, *songListItem)
	}

	response := web.SongListResponse{
		Songs: songList,
	}

	return response
}

func (service *SongServiceImpl) GetSongByArtistId(artistId string) web.SongArtistResponse {
	songList, err := service.SongRepository.GetAllUnitByArtistId(artistId)
	helper.PanicIfError(err)

	response := web.SongArtistResponse{
		Songs: songList,
	}

	return response
}

func (service *SongServiceImpl) getTopSimilarSong(songId int) []int {
	return service.getTopKSimilarSong(songId, 5)
}

type SongSimilarity struct {
	SongId     int
	Similarity float64
}

func sortSongSimilarity(similarities []SongSimilarity) {
	sort.Slice(similarities, func(i, j int) bool {
		return similarities[i].Similarity > similarities[j].Similarity
	})
}

func (service *SongServiceImpl) getTopKSimilarSong(songId int, k int) []int {
	recommendedSongIdList, ok := service.SongRecommendation[songId]
	if ok {
		return recommendedSongIdList
	}

	sameArtistSongIdList, err := service.SongRepository.GetAllSongIdFromSameArtist(songId)
	helper.PanicIfError(err)

	similarityCol := service.SongSimilarityData.Col(strconv.Itoa(songId))

	var songSimList []SongSimilarity
	var sameArtistSongSimList []SongSimilarity

	songIds := service.SongSimilarityData.Col("SongId").Records()
	for i, simValue := range similarityCol.Float() {
		songId, _ := strconv.Atoi(songIds[i])
		songSimilarity := SongSimilarity{
			SongId:     songId,
			Similarity: simValue,
		}

		if slices.Contains(sameArtistSongIdList, songSimilarity.SongId) {
			sameArtistSongSimList = append(sameArtistSongSimList, songSimilarity)
		} else {
			songSimList = append(songSimList, songSimilarity)
		}
	}

	var topKSongIdList []int

	// 60% of recommendations are preferred from the same artist
	sameArtistSongCount := int(math.Ceil(float64(k) * 0.0))

	if len(sameArtistSongIdList) <= sameArtistSongCount {
		topKSongIdList = append(topKSongIdList, sameArtistSongIdList...)
		k -= len(sameArtistSongIdList)
	} else {
		sortSongSimilarity(sameArtistSongSimList)

		for i := 0; i < sameArtistSongCount; i++ {
			topKSongIdList = append(topKSongIdList, sameArtistSongSimList[i].SongId)
		}
		k -= sameArtistSongCount
	}

	sortSongSimilarity(songSimList)

	topK := songSimList[1 : k+1] // exclude the song itself

	for _, song := range topK {
		similarSongId := song.SongId
		topKSongIdList = append(topKSongIdList, similarSongId)
	}

	service.SongRecommendation[songId] = topKSongIdList

	return topKSongIdList
}

func (service *SongServiceImpl) SearchSongByName(name string) web.SongListResponse {
	songUnitList, err := service.SongRepository.GetAllUnitByName(name)
	helper.PanicIfError(err)

	var songList []songModel.SongListItem
	for _, songUnit := range songUnitList {
		artist, err := service.ArtistRepository.GetAllUnitBySongId(songUnit.SongId)
		helper.PanicIfError(err)

		songListItem := songModel.NewSongListItem(songUnit, artist)
		songList = append(songList, *songListItem)
	}

	response := web.SongListResponse{
		Songs: songList,
	}

	return response
}
