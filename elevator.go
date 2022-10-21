package main

import "sort"

type Elevator struct {
	ID                    int
	status                string
	currentFloor          int
	direction             string
	door                  Door
	floorRequestsList     []int
	completedRequestsList []int
}

func NewElevator(_elevatorID int, _currentFloor int) *Elevator {
	elevator := &Elevator{
		ID:                    _elevatorID,
		status:                "idle",
		currentFloor:          _currentFloor,
		direction:             "any",
		door:                  *NewDoor(1),
		floorRequestsList:     []int{},
		completedRequestsList: []int{},
	}
	return elevator
}

/**
* Move elevator to requested floor
**/
func (e *Elevator) move() {
	for len(e.floorRequestsList) != 0 {
		e.status = "moving"
		e.sortFloorList()
		destination := e.floorRequestsList[0]
		if e.direction == "up" {
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.direction == "down" {
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.operateDoors()
		e.completedRequestsList = append(e.completedRequestsList, e.floorRequestsList[0])
		e.floorRequestsList = e.floorRequestsList[1:]
	}
	e.status = "idle"
}

/**
 * Manage doors
 **/
func (e *Elevator) operateDoors() {
	if e.status == "stopped" || e.status == "idle" {
		e.door.status = "open"
		if len(e.floorRequestsList) < 1 {
			e.direction = ""
			e.status = "idle"
		}
	}
}

/**
* Sort the list of floor requested according to elevator direction
**/
func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Ints(e.floorRequestsList)
	} else if e.direction == "down" {
		sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
	}
}

/**
 * User press a button outside the elevator
 **/
func (e *Elevator) addNewRequest(requestedFloor int) {
	if !contains(e.floorRequestsList, requestedFloor) {
		e.floorRequestsList = append(e.floorRequestsList, requestedFloor)
	}
	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > requestedFloor {
		e.direction = "down"
	}
}
