package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/Roverr/seat-distributor/core"

	"github.com/julienschmidt/httprouter"
)

// API describes which functions should be implemented by the business logic
type API interface {
	ListFlights(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	ReserveSeat(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	ListSeats(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// General implementation of the API
type General struct {
	Schedule *core.Schedule
}

// NewGeneralAPI creates a new instance of the general API implementation
func NewGeneralAPI(schedule *core.Schedule) API {
	return General{schedule}
}

// sendJSON is to send a given structure as JSON to the client
func (g General) sendJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		// Should log error here
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

// ListFlights lists all the available flights in the system
func (g General) ListFlights(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	values := r.URL.Query()
	limit, err := strconv.Atoi(values.Get("limit"))
	if err != nil {
		limit = 0
	}
	skip, err := strconv.Atoi(values.Get("skip"))
	if err != nil {
		skip = 0
	}
	var details []*core.Detail
	for i, detail := range g.Schedule.Flights {
		// Continue if skip is defined and we are not there yet
		if skip != 0 && i < skip {
			continue
		}
		// Break out of loop if we have the amount of data defined in limit
		if limit != 0 && len(details) == limit {
			break
		}
		details = append(details, detail)
	}
	detailsDto := map[string]interface{}{
		"flights": details,
		"limit":   limit,
		"skip":    skip,
		"count":   len(g.Schedule.Flights),
	}
	g.sendJSON(w, detailsDto)
}

// ListSeats is for listing available seats on a given flight
func (g General) ListSeats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	flight := g.Schedule.FindFlight(id)
	if flight == nil {
		w.WriteHeader(404)
		return
	}
	flightDto := map[string]interface{}{
		"seats": flight.Plane.GetFlightDTO(),
		"id":    id,
	}
	g.sendJSON(w, flightDto)
}

// ReserveSeat is for creating a reservation for a flight's seat (or multiple seats)
func (g General) ReserveSeat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// Should log error here
		w.WriteHeader(500)
		return
	}
	flight := g.Schedule.FindFlight(id)
	if flight == nil {
		// Should log error here
		w.WriteHeader(404)
		return
	}

	var reservation map[string][]string
	if err = json.Unmarshal(data, &reservation); err != nil {
		// Should log error here
		w.WriteHeader(500)
		return
	}
	reserves, ok := reservation["reserves"]
	if !ok {
		// Should log error here
		w.WriteHeader(400)
		return
	}
	var wg sync.WaitGroup
	var errors []error
	for _, shortcut := range reserves {
		wg.Add(1)
		go func(shortcut string) {
			err := flight.Plane.ReserveSeat(shortcut)
			if err != nil {
				errors = append(errors, err)
			}
			wg.Done()
		}(shortcut)
	}
	wg.Wait()
	if len(errors) == 0 {
		w.WriteHeader(200)
		return
	}

	errorsDto := map[string][]error{
		"errors": errors,
	}
	w.WriteHeader(500)
	g.sendJSON(w, errorsDto)
}
