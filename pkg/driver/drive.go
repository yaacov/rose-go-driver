package driver

import (
	"errors"

	"github.com/sirupsen/logrus"
)

var DriverName = "Go Cart"

// GameData represents the structure of the game's data, including the car's position
// and the track layout.
type GameData struct {
	Info struct {
		Car struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"car"`
	} `json:"info"`
	Track [][]string `json:"track"`
}

var forwardObstacleScores = map[string]int{
	obstacles.None:    0,
	obstacles.Penguin: 10,
	obstacles.Water:   4,
	obstacles.Crack:   5,
	obstacles.Trash:   -10,
	obstacles.Bike:    -10,
	obstacles.Barrier: -10,
}

var lrObstacleScores = map[string]int{
	obstacles.None:    0,
	obstacles.Penguin: 0,
	obstacles.Water:   -10,
	obstacles.Crack:   -10,
	obstacles.Trash:   -10,
	obstacles.Bike:    -10,
	obstacles.Barrier: -10,
}

var obstacleActionMap = map[string]string{
	obstacles.Penguin: actions.Pickup,
	obstacles.Crack:   actions.Jump,
	obstacles.Water:   actions.Brake,
}

func bestScoreForCell(y, x int, array [][]string, bestScores [][]int) int {
	possibleScores := []int{}

	// To directly above
	forwardObstacle := array[y-1][x]
	forwardPathScore := bestScores[y-1][x]
	possibleScores = append(possibleScores, forwardObstacleScores[forwardObstacle]+forwardPathScore)

	// To top-left
	if x > 0 {
		leftObstacle := array[y-1][x-1]
		leftPathScore := bestScores[y-1][x-1]
		possibleScores = append(possibleScores, lrObstacleScores[leftObstacle]+leftPathScore)
	}

	// To top-right
	if x < 2 {
		rightObstacle := array[y-1][x+1]
		rightPathScore := bestScores[y-1][x+1]
		possibleScores = append(possibleScores, lrObstacleScores[rightObstacle]+rightPathScore)
	}

	// Return the best score among the possible scores
	return max(possibleScores)
}

func calculateBestScores(array [][]string) [][]int {
	bestScores := make([][]int, 8)
	for i := range bestScores {
		bestScores[i] = make([]int, 3)
	}

	for y := 1; y < 8; y++ {
		for x := 0; x < 3; x++ {
			bestScores[y][x] = bestScoreForCell(y, x, array, bestScores)
		}
	}

	return bestScores
}

func bestAction(y, x int, array [][]string, bestScores [][]int) (string, error) {
	if y == 0 {
		return actions.None, nil
	}

	forwardObstacle := array[y-1][x]
	forwardScore := forwardObstacleScores[forwardObstacle] + bestScores[y-1][x]

	bestActScore := forwardScore
	bestAct := obstacleActionMap[forwardObstacle]

	// For moving left
	if x > 0 {
		leftObstacle := array[y-1][x-1]
		leftScore := lrObstacleScores[leftObstacle] + bestScores[y-1][x-1]
		if leftScore > bestActScore {
			bestActScore = leftScore
			bestAct = actions.Left
		}
	}

	// For moving right
	if x < 2 {
		rightObstacle := array[y-1][x+1]
		rightScore := lrObstacleScores[rightObstacle] + bestScores[y-1][x+1]
		if rightScore > bestActScore {
			bestActScore = rightScore
			bestAct = actions.Right
		}
	}

	return bestAct, nil
}

func extractArray(track [][]string, carX int) [][]string {
	// Determine the starting column based on carX
	startCol := 0
	if carX >= 3 {
		startCol = 3
	}

	// Extract the 3x8 slice from track
	array := make([][]string, 8)
	for i := 0; i < 8; i++ {
		array[i] = track[i][startCol : startCol+3]
	}

	return array
}

// Drive determines the next action for the car based on the current game data.
// It calculates the best scores for each cell in the track and then determines
// the best action to take from the car's current position.
//
// Parameters:
//   - gameData: The current state of the game, including the car's position and track layout.
//
// Returns:
//   - string: The chosen action for the car.
//   - error: An error object if there was an issue determining the best action.
func Drive(gameData GameData) (string, error) {
	car := gameData.Info.Car
	array := extractArray(gameData.Track, car.X)

	bestScores := calculateBestScores(array)

	action, err := bestAction(car.Y, car.X%3, array, bestScores)
	if err != nil {
		return actions.None, errors.New("Error determining best action")
	}

	if logrus.GetLevel() == logrus.DebugLevel {
		DebugPrint(gameData, bestScores, action)
	}

	return action, nil
}
