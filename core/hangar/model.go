package hangar

import (
	"fmt"
)

var (
	// ErrSeatAlreadyTakenFn is a function to create a specific error if a seat is already taken
	ErrSeatAlreadyTakenFn = func(shortcut string) error { return fmt.Errorf("%s is already taken", shortcut) }

	// ErrSeatNotExistsFn is a function to create a specific error if a seat does not exist
	ErrSeatNotExistsFn = func(shortcut string) error { return fmt.Errorf("%s does not exist", shortcut) }
)

// Seat indicates a single seat in a block along
// with the shortcut for it and if it is taken or not
type Seat struct {
	Taken    bool   `json:"taken"`
	Shortcut string `json:"shortcut"`
}

// Block is an aisle in a row with multiple seats
type Block struct {
	Seats []*Seat `json:"seats"`
}

// Row is a single row on a plane
type Row struct {
	Blocks []*Block `json:"blocks"`
}

// Plane is a struct that holds an array of rows that indicates
// how the plane would look like in a matrix
type Plane []*Row

// GetFlightDTO is to fetch all seats in an easily readable form
func (p Plane) GetFlightDTO() []*Seat {
	var seats []*Seat
	for _, row := range p {
		for _, block := range row.Blocks {
			for _, seat := range block.Seats {
				seats = append(seats, seat)
			}
		}
	}
	return seats
}

// ReserveSeat is to reserve a seat on a plane
func (p *Plane) ReserveSeat(shortcut string) error {
	for _, row := range *p {
		for _, block := range row.Blocks {
			for _, seat := range block.Seats {
				if seat.Shortcut != shortcut {
					continue
				}
				if seat.Taken {
					return ErrSeatAlreadyTakenFn(shortcut)
				}
				seat.Taken = true
				return nil
			}
		}
	}
	return ErrSeatNotExistsFn(shortcut)
}
