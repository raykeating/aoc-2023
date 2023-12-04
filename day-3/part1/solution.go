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

    var total_sum int = 0
    for i, _ := range grid {
        // fmt.Println(sum_row(grid, i))
        // fmt.Println(row)
        total_sum += sum_row(grid, i)
    }

    fmt.Println(total_sum)
    

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
}

// this function gets the set of numbers in a line and their start/end values
func sum_row(grid [][]string, row int) int {

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
                if (is_part_number(grid, row, startIndex, i, buffer)) {
                    num, _ := strconv.Atoi(buffer)
                    sum += num
                }
            }
        } else {
            prevCharWasDigit = false
            if (is_part_number(grid, row, startIndex, i, buffer)) {
                num, _ := strconv.Atoi(buffer)
                sum += num
            }
            buffer = ""
        }
        
    }

    return sum
}

// this function gets a border around a line and checks whether it contains any symbols
// this should be a string instead of a string array
func is_part_number(grid [][]string, row int, col_start int, col_end int, buffer string) bool {

    var border []string

    // fmt.Println(grid[row][col_start:col_end])
    
    // read all diagonals
    border = append(border, get_item(grid, row+1, col_start-1)) // bottom left
    border = append(border, get_item(grid, row+1, col_end)) // bottom right
    border = append(border, get_item(grid, row-1, col_start-1)) // top left
    border = append(border, get_item(grid, row-1, col_end)) // top right

    // read left and right
    border = append(border, get_item(grid, row, col_start-1))
    border = append(border, get_item(grid, row, col_end))

    // read above
    for i := 0; i <= col_end - col_start; i++ {
        border = append(border, get_item(grid, row-1, col_start+i))
    }

    // read below
    for i := 0; i <= col_end - col_start; i++ {
        border = append(border, get_item(grid, row+1, col_start+i))
    }

    // fmt.Println(border)

    for _, v := range border {
        r, _ := regexp.Compile(`[.1234567890]`)
        if (!r.MatchString(v)) {
            return true
        }
    }

    
    fmt.Println("buffer", buffer)
    fmt.Println("border", border)
    return false
}

func get_item(grid [][]string, row int, col int) string {
    if (row < 0 || col < 0) { return "." }
    if (col >= len(grid[0]) || row >= len(grid)) { return "." }
    return grid[row][col]
}

func is_digit(char string) bool {
    r, _ := regexp.Compile(`\d`)
    return r.MatchString(char)
}