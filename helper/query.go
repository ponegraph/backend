package helper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func ExecuteGraphDBQuery(query string) ([]byte, error) {
	endpoint := os.Getenv("GRAPHDB_BASE_ENDPOINT")
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

	reqURL := fmt.Sprintf("%s?query=%s", endpoint, encodedQuery)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, errors.New(ErrFailedDatabaseQuery)
	}

	// Set header
	req.Header.Set("Accept", "application/sparql-results+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(ErrFailedDatabaseQuery)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(ErrUnableToReadDatabase)
	}
	return body, nil
}
