package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)
var (
	maximumChar int
	words []string
	wordsSize []int
	filePath string
)

func wordWrap() {

	// First we calculate all possible extera values at the end of line contains
	// word i to j
	extras := make([][]int, len(words) + 1 )
	for i := range extras {
		extras[i] = make([]int, len(words) + 1)
	}
	for i := 1; i < len(extras); i++ {
		extras[i][i] = maximumChar - wordsSize[i - 1]
		for j := i + 1 ; j < len(extras); j++  {
			extras[i][j] = extras[i][j-1] - wordsSize[j- 1] - 1
		}
	}
	for i := range extras {
		fmt.Println(extras[i])
	}
	// Next Step is to find every line cost by extras which was calculated.
	lineCost := make([][]int, len(words) + 1 )
	for i := range extras {
		lineCost[i] = make([]int, len(words) + 1)
	}
	for i := 1; i < len(lineCost) ;i ++ {
		for j := i  ; j < len(lineCost); j++ {
			if extras[i][j] < 0 {
				lineCost[i][j] = math.MaxInt32
			}else if j == len(lineCost) - 1 {
				lineCost[i][j] = 0
			} else {
				lineCost[i][j] = (extras[i][j] - ( j - i) ) * (extras[i][j] - ( j - i) ) * (extras[i][j] - ( j - i) )
			}
		}
	}
	fmt.Println(lineCost)

	// n = len(words)
	// when now costs[n] = min( 1<=i<=n ) { costs[i - 1] + lineCost[i, n] } is our goal.
	// We Convert this recursive approach to DP approach.

	results  := make([]int, len(words) + 1)
	costs := make([]int, len(words) + 1)
	costs[0] = 0
	for j :=1; j < len(costs) ; j++ {
		costs[j] = math.MaxInt32
		for i := 1; i <= j; i++ {
			if costs[j - 1] != math.MaxInt32 && lineCost[i][j] != math.MaxInt32 {
				if costs[j] > costs[ i - 1 ] + lineCost[i][j] {
					costs[j]  = costs[ i - 1 ] + lineCost[i][j]
					results[j] = i
				}
			}
		}
	}
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("Failed to create output.txt")
	}
	fmt.Println(results)
	storeResults(results, len(results) - 1 , f)
	f.Close()
}

// storeResults recursively find lines and store them in new file called output.txt
func storeResults(results []int, end int, f *os.File)   {
	if results[end] == 0 {
		return
	}else {
		storeResults(results, results[end] - 1, f)
			var line []string
			line = words[results[end]  - 1 : end]
			for i, word := range line {
				if i == len(line)-1 {
					_, err := f.WriteString(word)
					if err != nil {
						log.Fatal("Failed to Write results")
					}
				} else {
					_, err := f.WriteString(word + " ")
					if err != nil {
						log.Fatal("Failed to Write results")
					}
				}
			}
			_, _ = f.WriteString("\n")

		return
	}

}

func main() {
	fmt.Println("Enter maximum character per line: ")
	_, err := fmt.Scanf("%d\n", &maximumChar)
	if err != nil {
		log.Fatalln("Bad input")
	}
fmt.Println("Enter path of file that contains words:  ")
	_, err =fmt.Scanf("%s\n", &filePath)
	if err != nil {
		log.Fatalln("Bad input")
	}

	readFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to Load File")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanWords)

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
		wordsSize = append(wordsSize, len(fileScanner.Text()))

	}

	err = readFile.Close()
	if err != nil {
		log.Fatal("Failed to close file...!")
	}

	wordWrap()
}
