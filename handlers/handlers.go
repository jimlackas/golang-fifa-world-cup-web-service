package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	year := req.URL.Query().Get("year")
	var winners []byte
	var err error

	if year == "" {
		winners, err = data.ListAllJSON()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		winners, err = data.ListAllByYear(year)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	_, err = res.Write(winners)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("X-ACCESS-TOKEN")
	isTokenValid := data.IsAccessTokenValid(accessToken)
	if !isTokenValid {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := data.AddNewWinner(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {

}
