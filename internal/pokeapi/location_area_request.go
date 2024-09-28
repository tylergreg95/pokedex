package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasRes := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasRes)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasRes, nil
}
