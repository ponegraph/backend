package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/ponegraph/backend/helper"
	"github.com/ponegraph/backend/model/web"
	"github.com/ponegraph/backend/service"
	"net/http"
	"strconv"
)

type SongControllerImpl struct {
	SongService service.SongService
}

func NewSongController(songService service.SongService) SongController {
	return &SongControllerImpl{
		SongService: songService,
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
