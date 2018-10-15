package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Roverr/seat-distributor/core/hangar"

	"github.com/Roverr/seat-distributor/core"

	"github.com/julienschmidt/httprouter"
)

func main() {
	schedule := createDevSchedule()
	api := NewGeneralAPI(schedule)

	router := httprouter.New()
	router.GET("/flights", api.ListFlights)
	router.GET("/flights/:id", api.ListSeats)
	router.POST("/flights/:id/reserve", api.ReserveSeat)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// createDevSchedule creates a new schedule with developer planes
func createDevSchedule() *core.Schedule {
	devPlane := "dev-plane"
	hangar := hangar.MemoryHangar{}
	hangar.ConstructPlane(devPlane, 60, 3, 4, 3)

	plane, err := hangar.LaunchPlane(devPlane)
	if err != nil {
		log.Fatal(err)
	}

	schedule := core.Schedule{
		Hangar: &hangar,
		Flights: []*core.Detail{
			&core.Detail{
				ID:      "1",
				TakeOff: time.Now().Add(time.Hour * 24),
				Plane:   plane,
				From:    "Budapest",
				To:      "New York",
				TypeID:  devPlane,
			},
			&core.Detail{
				ID:      "2",
				TakeOff: time.Now().Add(time.Hour * 48),
				Plane:   plane,
				From:    "New York",
				To:      "Budapest",
				TypeID:  devPlane,
			},
		},
	}

	return &schedule
}
