package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ArtistController interface {
	GetArtistDetail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Search(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
