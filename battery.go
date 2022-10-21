package main

import (
	"math"
	"strconv"
)

type Battery struct {
	columnID                  int
	ID                        int
	status                    string
	columnsList               []*Column
	floorRequestButtonsList   []*FloorRequestButton
	amountOfColumns           int
	amountOfFloors            int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	floorRequestButtonID      int
}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	battery := &Battery{
		ID:                        _id,
		status:                    "online",
		amountOfColumns:           _amountOfColumns,
		amountOfFloors:            _amountOfFloors,
		amountOfBasements:         _amountOfBasements,
		amountOfElevatorPerColumn: _amountOfElevatorPerColumn,
		columnID:                  'A',
		floorRequestButtonID:      1,
	}
	if _amountOfBasements > 0 {
		battery.createBasementFloorRequestButtons(_amountOfBasements)
		battery.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}
	battery.createFloorRequestButtons(_amountOfFloors)
	battery.createColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn)
	return battery
}

func (b *Battery) createBasementColumn(amountOfBasements int, amountOfElevatorPerColumn int) {
	servedFloorsList := []int{}
	floor := -1
	for i := 0; i < amountOfBasements; i++ {
		servedFloorsList = append(servedFloorsList, floor)
		floor--
	}
	column := NewColumn(strconv.Itoa(b.columnID), amountOfBasements, amountOfElevatorPerColumn, servedFloorsList, true)
	b.columnsList = append(b.columnsList, column)
	b.columnID++
}

func (b *Battery) createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfElevator int) {
	amountOfFloorsPerColumn := int(math.Ceil(float64(_amountOfFloors) / float64(_amountOfColumns)))
	floor := 1
	for c := 1; c <= _amountOfColumns; c++ {
		servedFloorsList := []int{}
		for i := 0; i < amountOfFloorsPerColumn; i++ {
			if floor <= _amountOfFloors {
				servedFloorsList = append(servedFloorsList, floor)
				floor++
			}
		}
		column := NewColumn(strconv.Itoa(b.columnID), _amountOfFloors, _amountOfElevator, servedFloorsList, false)
		b.columnsList = append(b.columnsList, column)
		b.columnID++
	}
}

func (b *Battery) createFloorRequestButtons(_amountOfFloors int) {
	for i := 0; i < b.amountOfFloors; i++ {
		floorRequestButton := NewFloorRequestButton(b.floorRequestButtonID, "up")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, floorRequestButton)
		b.floorRequestButtonID++
	}
}

func (b *Battery) createBasementFloorRequestButtons(amountOfBasements int) {
	buttonFloor := -1
	for i := 0; i < amountOfBasements; i++ {
		basementButton := NewFloorRequestButton(buttonFloor, "down")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, basementButton)
		buttonFloor--
	}
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	var bestColumn *Column = nil
	for _, column := range b.columnsList {
		if contains(column.servedFloorList, _requestedFloor) {
			bestColumn = column
		}
	}
	return bestColumn
}

/**
* Simulate when a user press a button at the lobby
**/
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := b.findBestColumn(_requestedFloor)
	elevator := column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return column, elevator
}
