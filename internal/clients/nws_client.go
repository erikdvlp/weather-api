package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-api/internal/models/nws"
)

func GetPoints(lat, lon float64) (*nws.Points, error) {
	url := fmt.Sprintf("https://api.weather.gov/points/%f,%f", lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get points data from NWS (%s)", resp.Status)
	}

	var points nws.Points
	if err := json.NewDecoder(resp.Body).Decode(&points); err != nil {
		return nil, err
	}

	return &points, nil
}

func GetGridPoints(office string, gridX, gridY int) (*nws.GridPoints, error) {
	url := fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d,%d/forecast", office, gridX, gridY)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get gridpoints data from NWS (%s)", resp.Status)
	}

	var gridPoints nws.GridPoints
	if err := json.NewDecoder(resp.Body).Decode(&gridPoints); err != nil {
		return nil, err
	}

	return &gridPoints, nil
}
