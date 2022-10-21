package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewCallButton(_floor int, _direction string) *CallButton {
	callButton := &CallButton{
		ID:        1,
		status:    "online",
		floor:     _floor,
		direction: _direction,
	}
	return callButton
}
