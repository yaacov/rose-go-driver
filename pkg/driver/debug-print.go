package driver

import "fmt"

// DebugPrint outputs the game's current state, including the track layout, best scores,
// car's position, and the chosen action for debugging purposes.
//
// Parameters:
//   - gameData: The current state of the game, including the car's position and track layout.
//   - bestScores: The calculated best scores for each cell in the track.
//   - action: The chosen action for the car.
func DebugPrint(gameData GameData, bestScores [][]int, action string) {
	car := gameData.Info.Car
	array := extractArray(gameData.Track, car.X)

	if action == "" {
		action = "none"
	}

	fmt.Println("Array (Track):")
	for _, row := range array {
		for _, cell := range row {
			fmt.Printf("| %-8s", cell)
		}
		fmt.Println("|")
	}

	fmt.Println("\nBest Scores:")
	for _, row := range bestScores {
		for _, score := range row {
			fmt.Printf("| %-3d", score)
		}
		fmt.Println("|")
	}

	fmt.Printf("\nCar Position: x=%d, y=%d", car.X, car.Y)
	fmt.Printf("\nChosen Action: %s\n", action)
}
