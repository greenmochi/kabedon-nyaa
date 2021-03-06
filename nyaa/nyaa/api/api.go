// Package api defines a simple API available to a client such as
// the gRPC service, testing purposes, or user simply running the program.
package api

import (
	"github.com/greenmochi/ultimate-nyaa/logger"
	"github.com/greenmochi/ultimate-nyaa/nyaa"
	"github.com/greenmochi/ultimate-nyaa/torrent"
)

// API TODO
type API struct {
	nyaa *nyaa.Client

	torrent *torrent.Client
}

// Setup TODO
func Setup() (*API, error) {
	torrentClient, err := torrent.NewClient()
	if err != nil {
		logger.Fatalf("unable to create new torrent client. err=%s\n", err)
		return nil, err
	}

	api := &API{
		nyaa: &nyaa.Client{
			URL: nyaa.DefaultURL(),
		},
		torrent: torrentClient,
	}
	return api, nil
}

// TearDown TODO
func (api *API) TearDown() {
	api.torrent.Close()
}

// Search TODO
func (api *API) Search(url *nyaa.URL) bool {
	api.nyaa.URL = url

	body, err := api.nyaa.Get()
	if err != nil {
		logger.Errorf("unable to perform GET request for query=%s\n", api.nyaa.URL.Query)
		return false
	}

	if ok := api.nyaa.Parse(body); !ok {
		logger.Errorf("unable to parse body for query=%s\n", api.nyaa.URL.Query)
		return false
	}

	return true
}

// GetCurrentResults TODO
func (api *API) GetCurrentResults() []nyaa.Result {
	return api.nyaa.Results
}
