package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	baseURl = "https://pokeapi.co/api/v2/"
)

type Client struct {
	httpClient http.Client
}

// New Client

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) ListLocations(pageURL *string) (ResponseLocations, error) {
	url := baseURl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocations{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return ResponseLocations{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("expected status %d received %d", http.StatusOK, response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return ResponseLocations{}, err
	}

	locationsResp := ResponseLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return ResponseLocations{}, err
	}

	return locationsResp, nil
}
