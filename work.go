package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	
	inputNums := []int64{}
	 scanner := bufio.NewScanner(os.Stdin)
	 for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
	 	if err != nil {
			panic(err)
		}
		inputNums = append(inputNums, num)
	}

}
func sort(n []int) {

	for i := 1; i < len(n); i++ {
		key := n[i]
		Is := i
		for j := i - 1; j > -1; j-- {
			if n[j] < key {
				break
			}
			n[j+1] = n[j]
			Is = j
		}
		n[Is] = key
		
	}
	
}
