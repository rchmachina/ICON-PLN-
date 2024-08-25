package main

import (
	"fmt"
	"math"
	"strings"

	"unicode"
)

// Fungsi  nomor 1
func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func nomor1(sentence string) string {

	words := strings.Fields(sentence)

	for i := 0; i < len(words); i++ {
		words[i] = reverseWord(words[i])
	}

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}


// fungsi nomorn 2
func nomor2() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print(i, " : FizzBuzz,")
		} else if i%3 == 0 {
			fmt.Print(i, ": Fizz,")
		} else if i%5 == 0 {
			fmt.Println(i, ": Buzz,")
		} else {
			continue
		}
	}
	
}

// fungsi nomor 3
func nomor3(n int) {

	fib := make([]int, n+1)

	
	fib[0] = 0
	if n > 0 {
		fib[1] = 1
	}

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
		fmt.Print(fib[i] , ", ")
	}

}

func nomor4(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit

}

// nomor 5
func nomor5(input []string) int {
	count := 0
	for _, char := range input {
		if len(char) == 1 && unicode.IsDigit(rune(char[0])) {
			count++
		}
	}
	return count
}

func main() {

	//nomor 1
	fmt.Println("\n \n nomor 1")
	bulkSentence := []string{"italem irad irigayaj", "iadab itsap ulalreb", "nalub kusutret gnalali"}
	for _, sentence := range bulkSentence {

		reversedSentence := nomor1(sentence)
		fmt.Println(reversedSentence)
	}

	//nomor 2
	fmt.Println("\n \n nomor 2")
	nomor2()
	//nomor 3
	fmt.Println("\n \n nomor 3")
	n := 10
	nomor3(n)


	// nomor 4
	fmt.Println("\n \n nomor 4")
	dataNomor4 := [][]int{
		{10, 9, 6, 5, 15},
		{7, 8, 3, 10, 8},
		{5, 12, 11, 12, 10},
		{7, 18, 27, 10, 29},
		{20, 17, 15, 14, 10},
	}

	// Iterasi melalui array 2D
	for i, prices := range dataNomor4 {
		fmt.Printf("max profit %d: %d\n", i+1, nomor4(prices))
	}

	fmt.Println("\n \n nomor 5")

	dataNomor5 := [][]string{
		{"2", "h", "6", "u", "y", "t", "7", "j", "y", "h", "8"},
		{"b", "7", "h", "6", "h", "k", "i", "5", "g", "7", "8"},
		{"7", "b", "8", "5", "6", "9", "n", "f", "y", "6", "9"},
		{"u", "h", "b", "n", "7", "6", "5", "1", "g", "7", "9"},
	}
	for _, input := range dataNomor5 {
		fmt.Println("Jumlah angka :", nomor5(input))
	}
}
