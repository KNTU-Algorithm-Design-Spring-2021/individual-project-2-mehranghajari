package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	dict map[string]int
)

func start(sentence string) []string{
	dp := make( map[string][]string)
	return wordBreak(sentence, dp)
}
func wordBreak(sentence string, dp map[string][]string) []string{

	if len(dp[sentence]) != 0{
		return dp[sentence]
	}
	var result []string
	if dict[sentence] == 1 {
		result = append(result, sentence)
}

	for i := 1 ; i< len(sentence); i++ {
		prefix := sentence[0: i ]
		if dict[prefix] == 1 {
			returnStringsList := wordBreak(sentence[i:], dp)

			for _, v :=range returnStringsList {
				result = append(result, prefix + " " + v)
			}
		}
	}
	dp[sentence] = result
	return result
}

func main() {
	filePath := "dict.txt"
	readFile , err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to read File")
	}
	f  := bufio.NewScanner(readFile)

	f.Split(bufio.ScanWords)
	dict = make(map[string]int)
	//dict["salam"] = 1
	//dict["hi"] = 1
	//dict["man"] = 1
	for f.Scan() {
		dict[f.Text()] = 1
	}
	fmt.Println(dict["abaca"] )
	sentence := "Whenthemadnessallaroundusstartstotakeit'stollTakeatripinsidemyheadeverybody's"
	for _, v := range start(sentence) {
		fmt.Println(v)
	}

}