package controller

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	"github.com/ponegraph/backend/helper"
	"github.com/ponegraph/backend/model/web"
	"github.com/ponegraph/backend/service"
)

type ArtistControllerImpl struct {
	ArtistService service.ArtistService
	logger        *slog.Logger
}

func NewArtistController(artistService service.ArtistService, logger *slog.Logger) ArtistController {
	return &ArtistControllerImpl{
		ArtistService: artistService,
		logger:        logger,
	}
}

func (controller *ArtistControllerImpl) GetArtistDetail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	artistId := params.ByName("artistId")
	artistDetailResponse := controller.ArtistService.GetArtistDetail(artistId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   artistDetailResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArtistControllerImpl) Search(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var artistListResponse web.ArtistListResponse

	tag := request.URL.Query().Get("tag")

	if tag != "" {
		tag, _ = url.QueryUnescape(tag)
		artistListResponse = controller.ArtistService.SearchArtistByTag(tag)
	} else {
		name := request.URL.Query().Get("name")
		name, _ = url.QueryUnescape(name)
		artistListResponse = controller.ArtistService.SearchArtistByName(name)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   artistListResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
