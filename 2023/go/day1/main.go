package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
    "strings"
)

func main() {
    input := []string{}
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        convertedLine := convertWordInt(line)
        input = append(input,convertedLine)
    }
    coords := []int{}
    for _, k := range(input) {
        coords = append(coords, getFirstLastDigit(k))
    }
    total := 0
    for _, k := range(coords) {
        total += k
    }
    fmt.Println(total)
}

func getFirstDigit(input string) string{
    var firstDigit string
    for _, k := range(input) {
        if unicode.IsDigit(k) {
            firstDigit = string(k)
            break
        }
    }
    return firstDigit
}

func getLastDigit(input string) string{
    var lastdigit rune
    for _, k := range(input) {
        if unicode.IsDigit(k) {
            lastdigit = k
        }
    }
    return string(lastdigit)
}

func getFirstLastDigit(input string) int{
    first := getFirstDigit(input)
    second := getLastDigit(input)
    num, _ := strconv.Atoi(first + second)
    return num
}


func convertWordInt(input string) string {
    words := map[string]string{
        "one":   "o1e",
        "two":   "t2o",
        "three": "t3e",
        "four":  "f4r",
        "five":  "f5e",
        "six":   "s6x",
        "seven": "s7n",
        "eight": "e8t",
        "nine":  "n9e",
    }
    for word, number := range(words){
        input = strings.Replace(input, word, number, -1)
    }
    return input
}
