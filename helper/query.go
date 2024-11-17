package helper

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func ExecuteGraphDBQuery(query string) ([]byte, error) {
	endpoint := os.Getenv("GRAPHDB_URL")
	return ExecuteSparqlQuery(query, endpoint)
}

func ExecuteDBpediaQuery(query string) ([]byte, error) {
	return ExecuteSparqlQuery(query, "https://dbpedia.org/sparql")
}

func ExecuteSparqlQuery(query string, endpoint string) ([]byte, error) {
	if endpoint == "" {
		return nil, errors.New("configuration error: Missing database endpoint")
	}

	encodedQuery := url.QueryEscape(query)

	reqURL := fmt.Sprintf("%s?query=%s", strings.TrimSpace(endpoint), encodedQuery)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		slog.Error("Failed to create request", "error", err.Error())
		return nil, errors.New(ErrFailedDatabaseQuery)
	}

	// Set header
	req.Header.Set("Accept", "application/sparql-results+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to execute request", "error", err.Error())
		return nil, errors.New(ErrFailedDatabaseQuery)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", "error", err.Error())
		return nil, errors.New(ErrUnableToReadDatabase)
	}
	return body, nil
}
