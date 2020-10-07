package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"

	"github.com/algorithms_illuminated/greedy_algorithms/huffman/huffman"
)

func main() {
	filePtr := flag.String("f", "tests/test1.txt", "filename")	
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	frequencies := []float64{}

	for scanner.Scan() {
		score, err := strconv.ParseFloat(strings.TrimSuffix(scanner.Text(), "\n"), 64)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)	
		}
		frequencies = append(frequencies, score)
	}

	tree := huffman.BuildTree(frequencies)
	treeBranchMin, treeBranchMax := tree.FindMinAndMaxLengthCodeWord()

	fmt.Println("treeBranchMin: ", treeBranchMin)
	fmt.Println("treeBranchMax: ", treeBranchMax)
}