package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func readFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var stringFile string
	for scanner.Scan() {
		stringFile += scanner.Text()
	}
	return stringFile

}

func getNumbers(file string) [][]int {
	var allNumbers [][]int
	doState := true
	for i := 0; i < len(file); i++ {
		if file[i] == 'd' && file[i+1] == 'o' &&
			file[i+2] == '(' && file[i+3] == ')' {
			doState = true
		} else if file[i] == 'd' && file[i+1] == 'o' &&
			file[i+2] == 'n' && file[i+3] == '\'' &&
			file[i+4] == 't' && file[i+5] == '(' && file[i+6] == ')' {
			doState = false
		}
		if doState {
			if file[i] == 'm' && file[i+1] == 'u' &&
				file[i+2] == 'l' && file[i+3] == '(' {
				var number string
				var numbers []int
				for j := i + 4; j < len(file); j++ {
					if unicode.IsDigit(rune(file[j])) {
						number += string(rune(file[j]))
					} else {
						if file[j] == ',' && len(numbers) == 0 {
							intNumber, err := strconv.Atoi(number)
							if err != nil {
								fmt.Println(err)
								return allNumbers
							}
							numbers = append(numbers, intNumber)
							number = ""
							continue
						} else if file[j] == ')' && len(numbers) != 0 {
							intNumber, err := strconv.Atoi(number)
							if err != nil {
								fmt.Println(err)
								return allNumbers
							}
							numbers = append(numbers, intNumber)
							allNumbers = append(allNumbers, numbers)
							break
						} else {
							break
						}
					}
				}
			}
		}
	}
	return allNumbers
}

func main() {
	file := readFile("day3.input")
	allNumbers := getNumbers(file)
	var multipliedNumbers []int
	for _, pair := range allNumbers {
		multipliedNumber := pair[0] * pair[1]
		multipliedNumbers = append(multipliedNumbers, multipliedNumber)
	}
	sumOfAll := 0
	for _, num := range multipliedNumbers {
		sumOfAll += num
	}
	fmt.Println(sumOfAll)
}
