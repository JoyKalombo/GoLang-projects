package main

import (
	"fmt"
	"os"
)

// This creates and empty slice with 9 lines and 9 columns.
var empty_sudoku = [9][9]int{}

func Formating() {
	// for every line in the map
	for _, line := range empty_sudoku {
		// for every square
		for index := range line {
			// print the value of each space of this line
			fmt.Print(line[index])
			// as long as we are not at the last value of this line
			if index != 8 {
				// add a space
				fmt.Print(" ")
			}
		}
		// change line when you print everything on this line
		fmt.Println()
	}
}

// returns if lines and columns and 3x3 boxes are not true
func IsValid(column int, line int, value int) bool {
	return !Vertical(column, line, value) &&
		!Horizontal(column, line, value) &&
		!Square3x3(column, line, value)
}

// returns if the value exists in this column
func Vertical(column int, line int, value int) bool {
	for index := range [9]int{} {
		if empty_sudoku[index][column] == value {
			return true
		}
	}
	return false
}

// returns if the value exists in this line
func Horizontal(column int, line int, value int) bool {
	for index := range [9]int{} {
		if empty_sudoku[line][index] == value {
			return true
		}
	}
	return false
}

// returns if the value exists in the 3x3 box
// the division returns the integer part only...
func Square3x3(column int, line int, value int) bool {
	sqr_column := int(column/3) * 3
	sqr_line := int(line/3) * 3
	for line_box := range [3]int{} {
		for column_box := range [3]int{} {
			if empty_sudoku[sqr_line+line_box][sqr_column+column_box] == value {
				return true
			}
		}
	}
	return false
}

// move to the next square
func Next(column int, line int) (int, int) {
	//
	square_in_line := (column + 1) % 9
	next_line := line
	// if you are at the end of the line
	if square_in_line == 0 {
		// change line
		next_line = line + 1
	}
	return square_in_line, next_line
}

// returns if this sudoku has a solution
func Solution(column int, line int) bool {
	// if we fill the last line there is a solution
	if line == 9 {
		return true
	}
	// if there is a sudoku number in this square
	if empty_sudoku[line][column] != 0 {
		// run again this function for the next square as defined from next funtion
		return Solution(Next(column, line))
	} else {
		// for every sudoku number
		for i := range [9]int{} {
			// add one
			value := i + 1
			// if this number does not violate any rules (line, columns box3x3)
			if IsValid(column, line, value) {
				// put that value in this square
				empty_sudoku[line][column] = value
				// if Solution() with the new parameters returns true return true
				if Solution(Next(column, line)) {
					return true
				}
				// back step a solution and try the new value
				empty_sudoku[line][column] = 0
			}
		}
		// if there is no solution at all return false
		return false
	}
}

func main() {
	// read the given arguments from the terminal
	arguments := os.Args
	// if the arguments are anything but 10
	if len(arguments) != 10 {
		// Print Error and terminate the program
		fmt.Println("Error")
		return
	}
	// for each of the arguments, while the arg are less than the length of the arguments, add one
	for arg := 1; arg < len(arguments); arg++ {
		// count the runes of each arg
		runes := []rune(arguments[arg])
		// if is not 9 Print Error and terminate
		if len(runes) != 9 {
			fmt.Println("Error")
			return
		}
		// for every number in a line
		for index := range runes {
			// if that num is a '.' replace it with 0
			if runes[index] == '.' {
				empty_sudoku[arg-1][index] = 0
			} else {
				// ascii 48 = 0. if the rune is a number and not a dot replace it from a rune to and int
				empty_sudoku[arg-1][index] = int(runes[index]) - 48
				// if this square is not a valid sudoku number Print Error and terminate
				if empty_sudoku[arg-1][index] < 1 || empty_sudoku[arg-1][index] > 9 {
					fmt.Println("Error")
					return
				}
			}
		}
	}

	// if there is a solution
	if Solution(0, 0) {
		// print it
		Formating()
		// else print error
	} else {
		fmt.Println("Error")
	}
}