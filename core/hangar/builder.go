package hangar

import (
	"fmt"
)

// nextChar is a function to figure out seat index
func nextChar(shift int) string {
	ch := byte('A')
	if ch += byte(shift); ch > 'Z' {
		return string('A')
	}
	return string(ch)
}

// constructSeats is for creating seats for a given row
func constructSeats(rowIndex, whichBlock, seatsInThisRow, numberOfSeats int) []*Seat {
	var seats []*Seat
	for i := 0; i < numberOfSeats; i++ {
		seats = append(seats, &Seat{
			Shortcut: fmt.Sprintf("%d%d%s", rowIndex, whichBlock, nextChar(seatsInThisRow+i)),
		})
	}
	return seats
}

// constructBlocks is for creating blocks based on given seats arrangement
func constructBlocks(rowIndex int, seats ...int) []*Block {
	blocks := []*Block{}
	// Keep tracking of how many seats in the row already
	appliedSeats := 0
	for _, numberOfSeats := range seats {
		blocks = append(blocks, &Block{
			Seats: constructSeats(rowIndex, len(blocks)+1, appliedSeats, numberOfSeats),
		})
		appliedSeats += numberOfSeats
	}
	return blocks
}

// constructPlane is to create a plane based on the number of rows and
// the arrangement
func constructPlane(numberOfRows int, arrangement ...int) []*Row {
	rows := []*Row{}
	for i := 0; i < numberOfRows; i++ {
		rows = append(rows, &Row{
			Blocks: constructBlocks(i+1, arrangement...),
		})
	}
	return rows
}
