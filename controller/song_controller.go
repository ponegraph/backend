package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type SongController interface {
	GetSongDetail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetTopRank(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Search(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
