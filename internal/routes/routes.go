package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"weather-api/internal/clients"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/health", getHealth)
	r.GET("/weather", getWeather)
}

func getHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "running",
	})
}

func getWeather(c *gin.Context) {
	// Validate input
	latStr := c.Query("lat")
	lonStr := c.Query("lon")
	lat, lon, err := validateCoordinates(latStr, lonStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get locational data from NWS
	points, err := clients.GetPoints(lat, lon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get forecast data from NWS
	office := points.Properties.GridID
	gridX := points.Properties.GridX
	gridY := points.Properties.GridY
	gridPoints, err := clients.GetGridPoints(office, gridX, gridY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return native data
	characterizeTemperature := func(temperature int) string {
		switch {
		case temperature < 32:
			return "cold"
		case temperature > 70:
			return "hot"
		default:
			return "moderate"
		}
	}
	shortForecast := gridPoints.Properties.Periods[0].ShortForecast
	temperature := gridPoints.Properties.Periods[0].Temperature
	characterization := characterizeTemperature(temperature)
	c.JSON(http.StatusOK, gin.H{
		"short_forecast": shortForecast,
		"temperature":    characterization,
	})
}

func validateCoordinates(latStr, lonStr string) (float64, float64, error) {
	validateCoordinate := func(s string, lower, upper int) (float64, error) {
		if s == "" {
			return 0, errors.New("coordinate not set")
		}
		coord, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse coordinate as float: %s", s)
		}
		if coord < float64(lower) || coord > float64(upper) {
			return 0, fmt.Errorf("coordinate not in valid range %d to %d: %s", lower, upper, s)
		}
		return coord, nil
	}

	lat, err := validateCoordinate(latStr, -90, 90)
	if err != nil {
		return 0, 0, fmt.Errorf("latitude validation failed: %s", err)
	}
	lon, err := validateCoordinate(lonStr, -180, 180)
	if err != nil {
		return 0, 0, fmt.Errorf("longitude validation failed: %s", err)
	}

	return lat, lon, nil
}
