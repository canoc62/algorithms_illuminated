package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"strconv"
)

type data struct {
	weight float64
	length float64
	score float64
}

type dataList []data

func (d dataList) Len() int {
	return len(d)
}

func (d dataList) Less(i, j int) bool {
	if d[i].score == d[j].score {
		return d[i].weight > d[j].weight
	}
	return d[i].score > d[j].score
}

func (d dataList) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	filePtr := flag.String("f", "test1.txt", "filename")	
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	events := dataList{}

	scanner.Scan()
	for scanner.Scan() {
		vals := strings.Split(strings.TrimSuffix(scanner.Text(), "\n"), " ")
		event := data{}
		w, err1 := strconv.ParseFloat(vals[0], 64)
		if err1 != nil {
			log.Fatalln(err1)
			os.Exit(1)	
		}
		event.weight = w

		l, err2 := strconv.ParseFloat(vals[1], 64) 
		if err2 != nil {
			log.Fatalln(err2)
			os.Exit(1)
		}
		event.length = l
		event.score = w/l
		events = append(events, event)
	}

	sort.Sort(events)
	fmt.Println(events)

	weightedTimeSum := events[0].weight*events[0].length
	for i := 1; i < len(events); i++ {
		events[i].length = events[i].length + events[i-1].length
		weightedTimeSum += events[i].weight*events[i].length
	}
	fmt.Println(events)

	fmt.Println("weightedSum", weightedTimeSum)
}