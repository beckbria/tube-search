package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Open questions:
//   * Route order?
//   * Zone number?

func fetchApi(url string) (string, error) {
	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(responseData), nil
}

func GetAllLines() ([]Tfl19, error) {
	var lines []Tfl19

	// Get a list of all tube routes.
	// https://api-portal.tfl.gov.uk/api-details#api=Line&operation=Line_RouteByModeByPathModesQueryServiceTypes
	apiRoutes, err := fetchApi("https://api.tfl.gov.uk/Line/Mode/tube/Route")
	if err != nil {
		return lines, err
	}

	json.Unmarshal([]byte(apiRoutes), &lines)
	return lines, nil
}

func GetStopsForLine(lineId string) ([]Tfl11, error) {
	var stops []Tfl11

	// Get a list of stops for a line
	// https://api-portal.tfl.gov.uk/api-details#api=Line&operation=Line_StopPointsByPathIdQueryTflOperatedNationalRailStationsOnly
	apiStops, err := fetchApi(fmt.Sprintf("https://api.tfl.gov.uk/Line/%s/StopPoints", lineId))
	if err != nil {
		return stops, err
	}

	json.Unmarshal([]byte(apiStops), &stops)
	return stops, nil
}

// Stop order for line
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Line/Line_RouteSequence
// https://api.tfl.gov.uk/Line/Line/{id}/Route/Sequence/{direction
