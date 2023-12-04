package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    var grid [][]string

    

    var i int = 0
    for scanner.Scan() {
        var line string = scanner.Text()
        grid = append(grid, strings.Split(line, ""))
        i++
    }

    // fmt.Println(is_part_number(grid, 0, 5, 7))

    // sum_row(grid, 4)

    var adjacent_to_stars []map[string]int

    for i, _ := range grid {
        // fmt.Println(sum_row(grid, i))
        // fmt.Println(row)
        sum_row(grid, i, adjacent_to_stars)
    }

    fmt.Println(adjacent_to_stars)

    
    

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
}

// this function gets the set of numbers in a line and their start/end values
func sum_row(grid [][]string, row int, adjacent_to_stars []map[string]int) int {

    var sum int = 0

    var buffer string = ""
    var prevCharWasDigit bool = false
    var startIndex int = -1

    for i, char := range grid[row] {
        if (is_digit(char)) {
            buffer += char
            fmt.Println(buffer)
            if (!prevCharWasDigit) {
                startIndex = i
                prevCharWasDigit = true
            }
            if (i == len(grid[row]) - 1) {
                get_adjacent_to_star(grid, row, startIndex, i, buffer, adjacent_to_stars)
            }
        } else {
            prevCharWasDigit = false
            get_adjacent_to_star(grid, row, startIndex, i, buffer, adjacent_to_stars)
            buffer = ""
        }
        
    }

    return sum
}

// this function gets a border around a line and checks whether it contains any symbols
// this should be a string instead of a string array
func get_adjacent_to_star(grid [][]string, row int, col_start int, col_end int, buffer string, adjacent_to_stars []map[string]int) {

    // fmt.Println(grid[row][col_start:col_end])
    
    // read all diagonals
    get_item(grid, row+1, col_start-1, buffer, adjacent_to_stars)
    get_item(grid, row+1, col_end, buffer, adjacent_to_stars)
    get_item(grid, row-1, col_start-1, buffer, adjacent_to_stars)
    get_item(grid, row-1, col_end, buffer, adjacent_to_stars)
    get_item(grid, row, col_start-1, buffer, adjacent_to_stars)
    get_item(grid, row, col_end, buffer, adjacent_to_stars)

    // read above & below
    for i := 0; i <= col_end - col_start; i++ {
        get_item(grid, row-1, col_start+i, buffer, adjacent_to_stars)
        get_item(grid, row+1, col_start+i, buffer, adjacent_to_stars)
    }
}

func get_item(grid [][]string, row int, col int, buffer string, adjacent_to_stars []map[string]int) {
    if (row < 0 || col < 0) { return }
    if (col >= len(grid[0]) || row >= len(grid)) { return }
    var item string = grid[row][col]
    if item == "*" {
        var adjacency_info map[string]int
        
        num, _ := strconv.Atoi(buffer)
        adjacency_info["number"] = num
        adjacency_info["row"] = num
        adjacency_info["col"] = num

        adjacent_to_stars = append(adjacent_to_stars, adjacency_info)
    }
}

func is_digit(char string) bool {
    r, _ := regexp.Compile(`\d`)
    return r.MatchString(char)
}