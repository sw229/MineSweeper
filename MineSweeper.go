package main

import (
	"fmt"
	"strconv"
)

func newField(mineNum int, field [][]int) [][]int {

	return field
}

func markCell(i, j int, visibleField [][]string) [][]string {
	visibleField[i][j] = "P"
	return visibleField
}

func openCell(i, j int, field [][]int, visibleField [][]string) (bool, [][]string) {
	if field[i][j] == -1 {
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

func main() {
	//var gameOver bool
	size_h := 10
	size_v := 10
	field := make([][]int, size_v)
	for i := range field {
		field[i] = make([]int, size_h)
	}
	field[1][1] = -1
	/*
	   for i := range(field){
	       for j := range(field[0]){
	           if field[i][j] != -1 {
	               for k := i - 1; k <= i + 1; k++{
	                   for l := j - 1; l <= j + 1; j++{
	                       if k < 0 || k > size_v - 1 || l < 0 || l > size_h - 1 {
	                           continue
	                       } else if field[k][l] == -1 {
	                           field[i][j] += 1
	                       }
	                   }
	               }
	           }
	       }
	   }
	*/

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

	//visibleField = markCell(1, 1, visibleField)
	_, visibleField = openCell(9, 9, field, visibleField)

	for i := range field {
		for j := range field[i] {
			fmt.Printf("%d ", field[i][j])
		}
		fmt.Println()
	}

	fmt.Println("\n\n")

	for i := range visibleField {
		for j := range visibleField[i] {
			fmt.Printf("%s ", visibleField[i][j])
		}
		fmt.Println()
	}

}
