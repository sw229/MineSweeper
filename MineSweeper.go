package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func newFields(mineNum, size_v, size_h int) ([][]int, [][]string) {
	field := make([][]int, size_v)
	for i := range field {
		field[i] = make([]int, size_h)
	}

	for i := 0; i < mineNum; i++ {
		x := rand.Intn(len(field))
		y := rand.Intn(len(field))
		if field[y][x] == -1 {
			i -= 1
		} else {
			field[y][x] = -1
		}
	}
	for i := range field {
		for j := range field[0] {
			if field[i][j] != -1 {
				if i > 0 && j > 0 && field[i-1][j-1] == -1 {
					field[i][j] += 1
				}
				if i > 0 && field[i-1][j] == -1 {
					field[i][j] += 1
				}
				if i > 0 && j < len(field[0])-1 && field[i-1][j+1] == -1 {
					field[i][j] += 1
				}
				if j > 0 && field[i][j-1] == -1 {
					field[i][j] += 1
				}
				if j < len(field[0])-1 && field[i][j+1] == -1 {
					field[i][j] += 1
				}
				if i < len(field)-1 && j > 0 && field[i+1][j-1] == -1 {
					field[i][j] += 1
				}
				if i < len(field)-1 && field[i+1][j] == -1 {
					field[i][j] += 1
				}
				if i < len(field)-1 && j < len(field[0])-1 && field[i+1][j+1] == -1 {
					field[i][j] += 1
				}
			}
		}
	}

	visibleField := make([][]string, size_v)
	for i := range visibleField {
		visibleField[i] = make([]string, size_h)
	}

	return field, visibleField
}

func markCell(i, j int, visibleField [][]string) [][]string {
	visibleField[i][j] = "P"
	return visibleField
}

func openCell(i, j int, field [][]int, visibleField [][]string) (bool, [][]string) {
	if visibleField[i][j] == "P" {
	} else if field[i][j] == -1 {
		return true, visibleField
	} else if field[i][j] != 0 {
		visibleField[i][j] = strconv.Itoa(field[i][j])
		return false, visibleField
	} else if visibleField[i][j] != "_" {
		visibleField[i][j] = "_"

		if i > 0 && j > 0 {
			_, visibleField = openCell(i-1, j-1, field, visibleField)
		}
		if i > 0 {
			_, visibleField = openCell(i-1, j, field, visibleField)
		}
		if i > 0 && j < len(field[0])-1 {
			_, visibleField = openCell(i-1, j+1, field, visibleField)
		}
		if j > 0 {
			_, visibleField = openCell(i, j-1, field, visibleField)
		}
		if j < len(field[0])-1 {
			_, visibleField = openCell(i, j+1, field, visibleField)
		}
		if i < len(field)-1 && j > 0 {
			_, visibleField = openCell(i+1, j-1, field, visibleField)
		}
		if i < len(field)-1 {
			_, visibleField = openCell(i+1, j, field, visibleField)
		}
		if i < len(field)-1 && j < len(field[0])-1 {
			_, visibleField = openCell(i+1, j+1, field, visibleField)
		}
	}

	return false, visibleField
}

func isVictory(field [][]int, visibleField [][]string) bool {
	for i := range field {
		for j := range field[0] {
			if field[i][j] != -1 && visibleField[i][j] != "_" {
				return false
			}
		}
	}
	return true
}

func main() {
	//var gameOver bool
	size_h := 4
	size_v := 4
	mineNum := 2
	gameOver := false

	field, visibleField := newFields(mineNum, size_v, size_h)

	for i := range field {
		for j := range field[i] {
			fmt.Printf("%d ", field[i][j])
		}
		fmt.Println()
	}

	fmt.Print("\n\n\n")

	for i := range visibleField {
		for j := range visibleField[i] {
			if visibleField[i][j] == "" {
				fmt.Print("O ")
			} else {
				fmt.Printf("%s ", visibleField[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Print("\n\n\n")

	for gameOver == false {
		var x, y int
		fmt.Println("input line")
		fmt.Scan(&y)
		fmt.Println("input column")
		fmt.Scan(&x)
		gameOver, visibleField = openCell(y, x, field, visibleField)
		gameOver = isVictory(field, visibleField)

		for i := range field {
			for j := range field[i] {
				fmt.Printf("%d ", field[i][j])
			}
			fmt.Println()
		}

		fmt.Print("\n\n\n")

		for i := range visibleField {
			for j := range visibleField[i] {
				if visibleField[i][j] == "" {
					fmt.Print("O ")
				} else {
					fmt.Printf("%s ", visibleField[i][j])
				}
			}
			fmt.Println()
		}

		fmt.Print("\n\n\n")
	}

}
