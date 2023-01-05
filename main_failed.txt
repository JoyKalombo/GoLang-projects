// package main

import (
	"fmt"
	"os"
)

func main() {
	starting_sudocu := os.Args[1:]
	PrintFormat(starting_sudocu)
	position := IsEmpty(starting_sudocu)
	line := position[0]
	column := position[1]
	if starting_sudocu[line][column] != '.' {
		IsEmpty(starting_sudocu)
	}

	// istrue := horizontal(num, position, starting_sudocu)

	try := 1
	fmt.Println(try, position)
	if !IsValidPosition(FindNext(try), position, starting_sudocu) && try <= 9 {
		fmt.Println(try, position)
		try++

	} else {
		// starting_sudocu[line][column] = FindNext(try)
		thisline := starting_sudocu[line]
		thisline[column] = FindNext(try)
	}
}

func PrintFormat(sudoku_map []string) {
	for _, row := range sudoku_map {
		fmt.Println(row)
		// fmt.Print('\n')
	}
}

func IsEmpty(sudocu []string) []int {
	// checks if the position is a dot
	position := []int{}
	for line := range sudocu {
		for column, sqr := range sudocu[line] {
			if sqr != '.' {
				// return line and return column (this is our position)
				continue
			} else {
				position = append(position, line, column)
			}
			break
		}
		break
	}
	return position
}

func IsValidPosition(num byte, position []int, sudocu []string) bool {
	// checks if the number can go in this possition
	fmt.Println(num)
	istrue := vertical(num, position, sudocu) &&
		horizontal(num, position, sudocu) &&
		littlegrid(num, position, sudocu)
	fmt.Println(istrue)

	return istrue
}

func horizontal(num byte, position []int, sudocu []string) bool {
	// checks if the number can go in the row... we run a for loop through each string
	// fmt.Println(num, position, "line1")
	line := position[0]
	// column := position[1]
	// fmt.Println(num, position, line, "line2")
	for i := range sudocu[line] {
		// fmt.Println(string(sudocu[line][i]), num, "line3")
		if sudocu[line][i] == num {
			// fmt.Println(string(sudocu[line][i]), num, "line4")
			return false
		}
	}
	// fmt.Println(num, position, "line5")
	return true
}

func vertical(num byte, position []int, sudocu []string) bool {
	// checks if the number can go in the column
	column := position[1]
	for i := range [9]int{} {
		if sudocu[i][column] == num {
			return false
		}
	}
	return true
}

func littlegrid(num byte, position []int, sudocu []string) bool {
	// 	// checks if the number can go in the little 3x3 grid
	line := position[0]
	column := position[1]
	sqr_line := int(line/3) * 3
	sqr_column := int(column/3) * 3
	for i_col := range [3]int{} {
		for i_line := range [3]int{} {
			if sudocu[sqr_column+i_col][sqr_line+i_line] == num {
				return false
			}
		}
	}
	return true
}

func CheckInput() {
	// checks if the input is valid
}

func FindNext(n int) byte {
	// returns a potential sollution for one position
	// str := ""
	numbers := "123456789"
	for num := n; num <= len(numbers); num++ {
		return byte(numbers[num])
	}
	return byte(numbers[0])
}

// go run . "39624...1" "17..6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
