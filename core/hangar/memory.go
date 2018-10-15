package hangar

import (
	"encoding/json"
)

// MemoryHangar is a struct to hold planes based on their names or IDs
type MemoryHangar map[string]Plane

// ConstructPlane is a function to store a new plane into the hangar
func (mh *MemoryHangar) ConstructPlane(id string, numberOfRows int, arrangement ...int) (Plane, error) {
	if mh == nil {
		return nil, ErrHangarNil
	}
	if plane, ok := (*mh)[id]; ok {
		return plane, ErrPlaneExists
	}
	plane := constructPlane(numberOfRows, arrangement...)
	(*mh)[id] = plane
	return plane, nil
}

// LaunchPlane is for creating a new plane instance from a plane in the hangar
func (mh *MemoryHangar) LaunchPlane(id string) (*Plane, error) {
	plane, ok := (*mh)[id]
	if !ok {
		return nil, nil
	}
	// Create a deep copy about the plane
	data, err := json.Marshal(plane)
	if err != nil {
		return nil, err
	}
	var copy Plane
	err = json.Unmarshal(data, &copy)
	return &copy, err
}

// ListPlanes is a function to list all available planes in the hangar by ID
func (mh *MemoryHangar) ListPlanes() map[string]Plane {
	return *mh
}
