// // Package main for the package containing the main() function
// package main

// We need the fmt library in order to read the user input and to display text in the console.
import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Declare the scenario struct, this is the main struct to hold everything.
type scenario struct {
	grid   coordinates
	rovers []rover
}

// This is the cooirdinates struct, holds X and Y axis number
type coordinates struct {
	xAxis int
	yAxis int
}

// This is the cooirdinates struct, hold the Rover ID, start position, movement instructions.
type rover struct {
	id           int
	start        position
	instructions string
}

// This is the position of the rover at the current time. Consists of an x axis, y axis and cardinal point to create one value.
type position struct {
	coordinates   coordinates
	cardinalPoint string
}

var (
	// //The input regex
	reGrid         string   = "test"
	rePosition     string   = "tter"
	reInstructions string   = "^[L|R|M]+$"
	cardinalPoint  []string = []string{"N", "E", "S", "W"}
)

// the main() function, this is where all golang projects start(just like most statically typed languages)
func main() {

	// Step one, Grab the grid limits.
	gridX, gridY := get_grid()
	fmt.Println(fmt.Sprintf("X axis: %d", gridX))
	fmt.Println(fmt.Sprintf("Y axis: %d", gridY))

	// Step three, loop through each rover to collect their starting point and driving instructions
	var rovers []rover

	for i := 1; i < (4 + 1); i++ {
		// Step 3.1 get the start location
		temp_start_loc := get_start_loc(i, gridX, gridY)
		fmt.Println(fmt.Sprintf("Rover #%d start coordinates: %s", i, temp_start_loc))
		// Step 3.2 get the instructions
		temp_directions := get_directions(i)
		fmt.Println(fmt.Sprintf("Rover #%d instructions: %s", i, temp_directions))

		t_rover := rover{id: i, instructions: temp_directions}
		rovers = append(rovers, t_rover)
	}

	// Now that we have all our rovers we need to process where they end up with the given starting location and instruction set
	// Assumptions:
	// 1. If two rovers end up on the same spot they can combine into a transformer.
	// 2. Rovers can't start on the same spot. We assume they all landed at the same time to save fuel.
	// 3. We are not going to log every stop, we just want the formula to get to the end point.
	fmt.Println(rovers)

}

func get_grid() (int, int) {
	var gridX int
	var gridY int

	// Step one, Grab the grid limits.
	// This is in a for loop so we can retry if invalid inputs are given.
	for gridX == 0 || gridY == 0 {
		fmt.Println("Please enter the grid size. Use a comma to seperate the x and y axis eg. '5,6'")
		var grid_input string //This will hold the input
		_, err := fmt.Scanln(&grid_input)
		if err != nil {
			fmt.Println(fmt.Sprintf("Input error: %s", err))
			continue
		}

		// Check for only one coma and giving relevant error message.
		if strings.Count(grid_input, ",") < 1 {
			fmt.Println(fmt.Sprintf("Input error: missing coma"))
			continue
		} else if strings.Count(grid_input, ",") > 1 {
			fmt.Println(fmt.Sprintf("Input error: too many comas"))
			continue
		}
		//validate the inputs are integers above the value of 0
		split_grid := strings.Split(grid_input, ",")
		gridX, err = strconv.Atoi(split_grid[0])
		if err != nil {
			fmt.Println(fmt.Sprintf("Input error: Invalid x axis value"))
			continue
		}
		gridY, err = strconv.Atoi(split_grid[1])
		if err != nil {
			fmt.Println(fmt.Sprintf("Input error: Invalid y axis value"))
			continue
		}

		// check the values are above min
		if gridX <= 0 {
			fmt.Println(fmt.Sprintf("Input error: X axis cannot be less than 1 "))
			continue
		}
		if gridY <= 0 {
			fmt.Println(fmt.Sprintf("Input error: Y axis cannot be less than 1 "))
			continue
		}

	}
	return gridX, gridY
}

func get_rover_count(max int) int {
	var count int
	for count == 0 {
		fmt.Println("Please enter the amount of rovers you are deploying")
		var rover_input string //This will hold the input
		_, err := fmt.Scanln(&rover_input)
		if err != nil {
			fmt.Println(fmt.Sprintf("Input error: %s", err))
			continue
		}

		tcount, err := strconv.Atoi(rover_input)
		if err != nil {
			fmt.Println(fmt.Sprintf("Input error: Invalid value, must be numeric"))
			continue
		}

		if tcount > max {
			fmt.Println(fmt.Sprintf("Input error: Invalid value, must be small or equal to 5"))
		}
		count = tcount
	}

	return count
}

func get_start_loc(id, gridX_limit, gridY_limit int) string {
	var valid bool
	var start_input string //This will hold the input

	// This is in a for loop so we can retry if invalid inputs are given.
	for !valid {
		fmt.Println(fmt.Sprintf("Please enter the starting location for Rover #%d", id))

		_, err := fmt.Scanln(&start_input)
		if err != nil {
			fmt.Println(fmt.Sprintf("Input error: %s", err))
			continue
		}

		// Check for 2 comas and give relevant error message.
		if !comma_count_valid(2, start_input) {
			continue
		}
		//validate the inputs are integers above the value of 0
		split_input := strings.Split(start_input, ",")

		// check the values are within range
		if !intMinMaxValid(split_input[0], 0, gridX_limit, "X axis") {
			continue
		}

		if !intMinMaxValid(split_input[1], 0, gridY_limit, "Y axis") {
			continue
		}

		if !stringInSlice(split_input[2], []string{"N", "E", "S", "W"}) {
			fmt.Println(fmt.Sprintf("Input error: Invalid cardinal direction given"))
			continue
		}
		valid = true
	}
	return start_input
}

func get_directions(id int) string {
	var valid bool
	var direction_input string //This will hold the input

	// This is in a for loop so we can retry if invalid inputs are given.
	for !valid {
		fmt.Println(fmt.Sprintf("Please enter the instructions for navigating Rover #%d", id))

		_, err := fmt.Scanln(&direction_input)
		if err != nil {
			fmt.Println(fmt.Sprintf("Input error: %s", err))
			continue
		}

		// re := regexp.MustCompile(`(?m)^[L|R|M]+$`)
		valid, err = regexp.MatchString("^[L|R|M]+$", direction_input)
		if err != nil || !valid {
			fmt.Println(fmt.Sprintf("Invalid format, only 'L', 'R', and 'M' are accepted"))
			continue
		}
	}
	return direction_input
}

//  ---- Tiny Helper methods ----
// These methods are used on multiple locations or are very basic methods to keep things clean.

func comma_count_valid(expected int, input_string string) bool {
	if strings.Count(input_string, ",") != expected {
		fmt.Println(fmt.Sprintf("Input error: invalid coma count, expected %d", expected))
		return false
	}
	return true
}

// method to see if a string is inside a slice.
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func intMinMaxValid(num string, min, max int, item string) bool {
	inum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(fmt.Sprintf("Input error: Invalid %s value", item))
	}
	if inum < min {
		fmt.Println(fmt.Sprintf("Input error: %s cannot be less than %d ", item, min))
		return false
	}
	if inum > max {
		fmt.Println(fmt.Sprintf("Input error: %s cannot be more than %d ", item, max))
		return false
	}
	return true
}
