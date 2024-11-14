package main

import (
	"github.com/ponegraph/backend/controller"
	"github.com/ponegraph/backend/helper"
	"github.com/ponegraph/backend/repository"
	"github.com/ponegraph/backend/service"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	os.Setenv("GRAPHDB_BASE_ENDPOINT", "http://localhost:7200/repositories/ponegraph-music")

	songRepository := repository.NewSongRepository()
	artistRepository := repository.NewArtistRepository()

	songService := service.NewSongService(songRepository, artistRepository)
	artistService := service.NewArtistService(songRepository, artistRepository)

	songController := controller.NewSongController(songService)
	artistController := controller.NewArtistController(artistService)

	router := httprouter.New()

	router.GET("/songs/id/:songId", songController.GetSongDetail)
	router.GET("/songs/top-rank", songController.GetTopRank)
	router.GET("/songs/search", songController.Search)

	router.GET("/artists/id/:artistId", artistController.GetArtistDetail)
	router.GET("/artists/search", artistController.Search)

	router.PanicHandler = controller.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
