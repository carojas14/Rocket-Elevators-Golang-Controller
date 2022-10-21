package main

type Door struct {
	ID     int
	status string
}

func NewDoor(id int) *Door {
	door := &Door{
		ID:     1,
		status: "closed",
	}
	return door
}
