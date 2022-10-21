package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID             int
	status         string
	requestedFloor int
	direction      string
}

func NewFloorRequestButton(_floor int, _direction string) *FloorRequestButton {
	floorRequestButton := &FloorRequestButton{
		ID:             1,
		status:         "online",
		requestedFloor: _floor,
		direction:      _direction,
	}
	return floorRequestButton
}
