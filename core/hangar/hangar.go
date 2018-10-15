package hangar

import (
	"errors"
)

var (
	// ErrNotFound is an error indicating that the given plane is not found
	ErrNotFound = errors.New("Plane not found")
	// ErrPlaneExists is an error indicating that the given plane already exists
	ErrPlaneExists = errors.New("Plane already exists")
	// ErrHangarNil is an error indicating that the hangar is nil
	ErrHangarNil = errors.New("Hangar has not been initalised")
)

// Hangar describes the business logic about how we handle
// planes in the system
type Hangar interface {
	ConstructPlane(id string, numberOfRows int, arrangement ...int) (Plane, error)
	LaunchPlane(id string) (*Plane, error)
	ListPlanes() map[string]Plane
}
