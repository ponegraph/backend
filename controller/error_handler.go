package controller

import (
	"errors"
	"fmt"
	"github.com/ponegraph/backend/exception"
	"github.com/ponegraph/backend/helper"
	"github.com/ponegraph/backend/model/web"
	"net/http"
)

type errorData struct {
	Message string `json:"message"`
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if actualErr, ok := err.(error); ok {
		if notFoundError(writer, request, actualErr) {
			return
		}
		if badRequestError(writer, request, actualErr) {
			return
		}
		internalServerError(writer, request, actualErr)
	} else {
		internalServerError(writer, request, errors.New(fmt.Sprintf("%v", err)))
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err error) bool {
	var notFoundErr *exception.NotFoundError
	if errors.As(err, &notFoundErr) {
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   errorData{notFoundErr.Error()},
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	}
	return false
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err error) bool {
	var badRequestErr *exception.BadRequestError
	if errors.As(err, &badRequestErr) {
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errorData{badRequestErr.Error()},
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err error) {
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   errorData{err.Error()},
	}

	helper.WriteToResponseBody(writer, webResponse)
}
