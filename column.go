package main

var elevatorID int = 1

type Column struct {
	ID                string
	status            string
	amountOfFloors    int
	amountOfElevators int
	elevatorsList     []*Elevator
	callButtonList    []*CallButton
	servedFloorList   []int
	isBasement        bool
}

func NewColumn(_id string, _amountOfFloors int, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	column := &Column{
		ID:                _id,
		status:            "online",
		amountOfElevators: _amountOfElevators,
		servedFloorList:   _servedFloors,
		isBasement:        _isBasement,
		elevatorsList:     []*Elevator{},
		callButtonList:    []*CallButton{},
	}
	column.createElevators(_amountOfFloors, _amountOfElevators)
	column.createCallButtons(_amountOfFloors, _isBasement)
	return column
}

/**
 * Create a list of call buttons for each column using CallButton class
 **/
func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	if c.isBasement {
		buttonFloor := -1
		for i := 0; i < len(c.servedFloorList); i++ {
			callButton := NewCallButton(buttonFloor, "up")
			c.callButtonList = append(c.callButtonList, callButton)
			buttonFloor--
		}
	} else {
		buttonFloor := 1
		for i := 0; i < len(c.servedFloorList); i++ {
			callButton := NewCallButton(buttonFloor, "down")
			c.callButtonList = append(c.callButtonList, callButton)
			buttonFloor++
		}
	}
}

/**
 * Create a list of elevators for each column using Elevator class
 **/
func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) {
	for i := 0; i < c.amountOfElevators; i++ {
		elevator := NewElevator(elevatorID, 1)
		c.elevatorsList = append(c.elevatorsList, elevator)
		elevatorID++
	}
}

/**
* Simulate when a user press a button on a floor to go back to the first floor
**/
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	elevator := c.findElevator(_requestedFloor, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	elevator.addNewRequest(1)
	elevator.move()
	return elevator
}

/**
* Find the best elevator, prioritizing the elevator that is already moving,
* that is closer to the user's floor and that goes to the same direction that user wants
**/
func (c *Column) findElevator(requestedFloor int, requestedDirection string) *Elevator {

	var bestElevator *Elevator
	var bestScore int = 6
	var referenceGap int = 10000000

	//Rquested floor is the lobby
	if requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {

			//The elevator is at the lobby and already has some requests. It is about to leave but has not yet departed
			if elevator.currentFloor == 1 && elevator.status == "stopped" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is at the lobby and has no requests
			} else if elevator.currentFloor == 1 && elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is lower than user and is coming up. It means that user is requesting an elevator to go to a basement, and the elevator is on same way to user.
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is above user and is coming down. It means that user is requesting an elevator to go to a floor, and the elevator is on same way to user
			} else if 1 < elevator.currentFloor && elevator.direction == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is not at the first floor, but doesn't have any request
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is not available, but still could take the call if nothing better is found
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			}
		}
		//Rquested floor is not lobby
	} else {
		for _, elevator := range c.elevatorsList {

			//The elevator is at the same level as user, and is about to depart to the first floor
			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is lower than user and is going up. User is on a basement, and the elevator can pick user up on it's way
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is higher than user and is going down. User is on a floor, and the elevator can pick user up on it's way
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is idle and has no requests
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is not available, but still could take the call if nothing better is found
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			}

		}
	}
	return bestElevator
}

/**
* Called by findElevator to compare current elevator in elevatorList with
* other elevators and return best elevator.
**/
func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator *Elevator, bestScore int, referenceGap int, bestElevator *Elevator, floor int) (*Elevator, int, int) {
	if scoreToCheck < bestScore {
		bestScore = scoreToCheck
		bestElevator = newElevator
		referenceGap = Abs(newElevator.currentFloor - floor)
	} else if bestScore == scoreToCheck {
		gap := Abs(newElevator.currentFloor - floor)
		if referenceGap > gap {
			bestElevator = newElevator
			referenceGap = gap
		}
	}
	return bestElevator, bestScore, referenceGap
}
