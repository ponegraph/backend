package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ponegraph/backend/controller"
	"github.com/ponegraph/backend/helper"
	"github.com/ponegraph/backend/repository"
	"github.com/ponegraph/backend/service"

	"github.com/julienschmidt/httprouter"
)

func CORSMiddleware(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		router.ServeHTTP(w, r)
	})
}

// LoggingMiddleware function that wraps the Router's ServeHTTP method
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Serve the request
		next.ServeHTTP(w, r)

		// Log request details after serving the request
		duration := time.Since(start).Seconds() * 1000
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // nil writer defaults to stdout
		logger.Info(
			"Request Info",
			slog.String("method", r.Method),
			slog.String("path", r.URL.String()),
			slog.String("duration", fmt.Sprintf("%.2fms", duration)),
			slog.String("client_ip", r.RemoteAddr),
		)
	})
}

func makeLog(w io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		AddSource: true,
		Level:     &level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.SourceKey:
				if source, ok := a.Value.Any().(*slog.Source); ok {
					source.File = filepath.Base(source.File)
				}
			case slog.TimeKey:
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}))
}

func main() {
	logger := makeLog(os.Stderr, slog.LevelInfo)
	songRepository := repository.NewSongRepository(logger)
	artistRepository := repository.NewArtistRepository(logger)

	songService := service.NewSongService(songRepository, artistRepository, logger)
	artistService := service.NewArtistService(songRepository, artistRepository, logger)

	songController := controller.NewSongController(songService, logger)
	artistController := controller.NewArtistController(artistService, logger)

	router := httprouter.New()

	router.GET("/songs/id/:songId", songController.GetSongDetail)
	router.GET("/songs/top-rank", songController.GetTopRank)
	router.GET("/songs/search", songController.Search)

	router.GET("/artists/id/:artistId", artistController.GetArtistDetail)
	router.GET("/artists/search", artistController.Search)

	router.PanicHandler = controller.ErrorHandler

	loggedRouter := CORSMiddleware(LoggingMiddleware(router))

	server := http.Server{
		Addr:    ":3000",
		Handler: loggedRouter,
	}

	logger.Info("Server started at localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
