package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter numbers: ")

	inputNums := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		inputNums = append(inputNums)
	}
	for i := 1; i < len(inputNums); i++ {
		x := inputNums[i]
		j := i
		for ; j >= 1 && inputNums[j-1] > x; j-- {
			inputNums[j] = inputNums[j-1]
		}
		inputNums[j] = x
		fmt.Println(inputNums)

	}
}
