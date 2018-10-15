package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/Roverr/seat-distributor/core/hangar"
)

var (
	// ErrSeatAlreadyReservedFn is a function to create a new specific error
	// for already taken seats
	ErrSeatAlreadyReservedFn = func(id, shortcut string) error { return fmt.Errorf("%s is already taken on plane %s", shortcut, id) }
)

// Schedule holds basic information about which plane takes of
// from, to, and when
type Schedule struct {
	Flights []*Detail
	Hangar  hangar.Hangar
}

// FindFlight is for getting an available flight from the list of flights
func (s *Schedule) FindFlight(id string) *Detail {
	for _, detail := range s.Flights {
		if detail.ID == id {
			return detail
		}
	}
	return nil
}

// ReserveSeat is for reserving a seat in a given plane
func (s *Schedule) ReserveSeat(id, shortcut string) error {
	flight := s.FindFlight(id)
	flight.mux.Lock()
	defer flight.mux.Unlock()
	return flight.Plane.ReserveSeat(shortcut)
}

// Detail describes information about the flight of a plane
type Detail struct {
	ID      string        `json:"id"`
	TakeOff time.Time     `json:"takeOff"`
	To      string        `json:"to"`
	From    string        `json:"from"`
	Plane   *hangar.Plane `json:"-"`
	TypeID  string        `json:"typeId"`
	mux     *sync.Mutex
}
