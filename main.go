package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Seat ...
type Seat struct {
	Taken    bool
	Shortcut string
}

// Block ...
type Block struct {
	Seats []*Seat
}

// Row ...
type Row struct {
	Blocks []*Block
}

func nextChar(shift int) string {
	ch := byte('A')
	if ch += byte(shift); ch > 'Z' {
		return string('A')
	}
	return string(ch)
}

func constructSeats(rowIndex, seatsInThisRow, numberOfSeats int) []*Seat {
	var seats []*Seat
	for i := 0; i < numberOfSeats; i++ {
		seats = append(seats, &Seat{
			Shortcut: fmt.Sprintf("%d%d%s", rowIndex, i+1, nextChar(seatsInThisRow+i)),
		})
	}
	return seats
}

func constructBlocks(rowIndex int, seats ...int) []*Block {
	blocks := []*Block{}
	appliedSeats := 0
	for _, numberOfSeats := range seats {
		blocks = append(blocks, &Block{
			Seats: constructSeats(rowIndex, appliedSeats, numberOfSeats),
		})
		appliedSeats += numberOfSeats
	}
	return blocks
}

// ListPlanes ...
func ListPlanes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func getDeveloperPlane() []*Row {
	rows := []*Row{}
	for i := 0; i < 2; i++ {
		rows = append(rows, &Row{
			Blocks: constructBlocks(i+1, 3, 4, 3),
		})
	}
	return rows
}

func main() {
	router := httprouter.New()
	router.GET("/planes", ListPlanes)

	log.Fatal(http.ListenAndServe(":8080", router))
}
