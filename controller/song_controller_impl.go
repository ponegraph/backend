package controller

import (
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/ponegraph/backend/helper"
	"github.com/ponegraph/backend/model/web"
	"github.com/ponegraph/backend/service"
)

type SongControllerImpl struct {
	SongService service.SongService
	logger      *slog.Logger
}

func NewSongController(songService service.SongService, logger *slog.Logger) SongController {
	return &SongControllerImpl{
		SongService: songService,
		logger:      logger,
	}
}

func (controller *SongControllerImpl) GetSongDetail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	songId := params.ByName("songId")
	songIdInt, _ := strconv.Atoi(songId)

	songDetail := controller.SongService.GetSongDetail(songIdInt)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   songDetail,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SongControllerImpl) GetTopRank(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	topRankSongResponse := controller.SongService.GetTopRank()

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   topRankSongResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SongControllerImpl) Search(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	name := request.URL.Query().Get("name")
	name, _ = url.QueryUnescape(name)

	songListResponse := controller.SongService.SearchSongByName(name)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   songListResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
