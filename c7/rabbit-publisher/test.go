// You can edit this code!
// Click here and start typing.
package test

import (
	"errors"
	"fmt"
	"strconv"
)

var limitX int
var limitY int
var lostRobots map[string]bool

// Representation of a robot
type Robot struct {
	// Whether robot is lost or not
	Lost bool
	// Where the robot is pointing to (N,E,S,W)
	Orientation string
	// Latitude coordinate representation of robot position
	X int
	// Longitude coordinate representation of robot position
	Y int
}

func (robot *Robot) GetPos() string {
	return strconv.Itoa(robot.X) + " " + strconv.Itoa(robot.Y)
}

func (robot *Robot) Instruct(direction string) (bool, error) {
	switch direction {
	case "F":
		return robot.Move()
	case "L":
		return robot.TurnLeft()
	case "R":
		return robot.TurnRight()
	default:
		return false, errors.New("Invalid instruction")
	}
}

// NOTE: There are smarter ways to do turn left and right, but the simplest possible (and most readable)
// is these two methods (two separate methods as code is not really being that reused here

// Function to turn right
func (robot *Robot) TurnRight() (bool, error) {
	switch robot.Orientation {
	case "N":
		robot.Orientation = "E"
		return false, nil
	case "E":
		robot.Orientation = "S"
		return false, nil
	case "W":
		robot.Orientation = "N"
		return false, nil
	case "S":
		robot.Orientation = "W"
		return false, nil
	default:
		return true, errors.New("Invalid orientation")
	}
}

// Function to turn left
func (robot *Robot) TurnLeft() (bool, error) {
	switch robot.Orientation {
	case "N":
		robot.Orientation = "W"
		return false, nil
	case "E":
		robot.Orientation = "N"
		return false, nil
	case "W":
		robot.Orientation = "S"
		return false, nil
	case "S":
		robot.Orientation = "E"
		return false, nil
	default:
		return true, errors.New("Invalid orientation")
	}
}

func (robot *Robot) Move() (bool, error) {
	switch robot.Orientation {
	case "N":
		if checkLostRobot(robot.X, robot.Y+1) {
			return false, nil
		}
		robot.Y++
		return robot.checkIsLost(), nil
	case "E":
		if checkLostRobot(robot.X+1, robot.Y) {
			return false, nil
		}
		robot.X++
		return robot.checkIsLost(), nil
	case "W":
		if checkLostRobot(robot.X-1, robot.Y) {
			return false, nil
		}
		robot.X--
		return robot.checkIsLost(), nil
	case "S":
		if checkLostRobot(robot.X, robot.Y-1) {
			return false, nil
		}
		robot.Y--
		return robot.checkIsLost(), nil
	default:
		return true, errors.New("Invalid orientation")
	}
}

func (robot *Robot) checkIsLost() bool {
	if robot.Y < 0 || robot.X < 0 || robot.Y >= limitY || robot.X >= limitX {
		robot.Lost = true
		lostRobots[robot.GetPos()] = true
		return true
	}
	return false
}

func checkLostRobot(x int, y int) bool {
	pos := strconv.Itoa(x) + " " + strconv.Itoa(y)
	if val, ok := lostRobots[pos]; ok {
		return val
	}
	return false
}

func main() {
	lostRobots = make(map[string]bool)

	limitX = 5
	limitY = 5

	chars := []rune("LRFFRFLF")
	robot := Robot{X: 0, Y: 0, Orientation: "N", Lost: false}
	gotLost := false
	for i := 0; i < len(chars) && !robot.Lost && !gotLost; i++ {
		char := string(chars[i])
		var err error
		gotLost, err = robot.Instruct(char)
		if err != nil {
			fmt.Println(err)
			gotLost = true
		}
	}
}
