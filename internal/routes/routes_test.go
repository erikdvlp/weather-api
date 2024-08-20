package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestValidateCoordinates(t *testing.T) {
	tests := []struct {
		latInput    string
		lonInput    string
		latExpected float64
		lonExpected float64
		errExpected bool
	}{
		{"27.950575", "-82.457176", 27.950575, -82.457176, false},
		{"90.0", "180.0", 90.0, 180.0, false},
		{"-90.0", "-180.0", -90.0, -180.0, false},
		{"90.1", "-82.457176", 0, 0, true},
		{"-90.1", "-82.457176", 0, 0, true},
		{"27.950575", "-180.1", 0, 0, true},
		{"invalid", "-82.457176", 0, 0, true},
		{"", "-82.457176", 0, 0, true},
		{"27.950575", "invalid", 0, 0, true},
		{"27.950575", "", 0, 0, true},
	}

	for _, test := range tests {
		latActual, lonActual, err := validateCoordinates(test.latInput, test.lonInput)
		if latActual != test.latExpected || lonActual != test.lonExpected || (err != nil) != test.errExpected {
			t.Errorf("expected %f %f %t; actual %f %f %t", test.latExpected, test.lonExpected, test.errExpected, latActual, lonActual, err != nil)
		}
	}
}

func TestGetWeather(t *testing.T) {
	tests := []struct {
		latInput     string
		lonInput     string
		codeExpected int
	}{
		{"27.950575", "-82.457176", 200},
		{"40.71427", "-74.00597", 200},
		{"32.78306", "-96.80667", 200},
	}

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	InitRoutes(r)

	for _, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/weather?lat=%s&lon=%s", test.latInput, test.lonInput), nil)
		r.ServeHTTP(w, req)

		if w.Code != test.codeExpected {
			t.Errorf("received unexpected response (%d): %s", w.Code, w.Body.String())
		}
	}
}
