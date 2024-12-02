package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	EntryFile := os.Args[1]
	Function := os.Args[2]

	if EntryFile == "" {
		return
	}

	if Function == "" {
		return
	}
	fu, err := strconv.Atoi(Function)
	if err != nil {
		panic(err)
	}

	switch fu {
	case 1:
		PartOne(EntryFile)
	case 2:
		PartTwo(EntryFile)
	}

	//fmt.Println(listArray)

}

func PartTwo(EntryFile string) {

	CounterMap := make(map[int]int)
	L1 := []int{}
	TotalSum := 0

	file, err := os.ReadFile(EntryFile)
	if err != nil {
		panic(err)
	}

	completeList := string(file)
	listArray := strings.Split(completeList, "\n")
	listArray = listArray[:len(listArray)-1]
	for _, value := range listArray {
		splitList := strings.Split(value, "   ")
		location, err := strconv.Atoi(splitList[0])
		if err != nil {
			panic(err)
		}

		_, found := CounterMap[location]
		if !found {
			CounterMap[location] = 0
		}
		L1 = append(L1, location)

		location, err = strconv.Atoi(splitList[1])
		if err != nil {
			panic(err)
		}

		value, found := CounterMap[location]
		if !found {
			CounterMap[location] = 1
		} else {
			CounterMap[location] = value + 1
		}

	}

	for _, value := range L1 {
		fmt.Printf("Index: %d, Value: %d\n", value, CounterMap[value])
		TotalSum += (value * CounterMap[value])
	}

	fmt.Printf("Total Sum: %d\n", TotalSum)

}

func PartOne(EntryFile string) {
	ListOne := &IntHeap{}
	ListTwo := &IntHeap{}
	TotalDist := 0.0

	heap.Init(ListOne)
	heap.Init(ListTwo)

	file, err := os.ReadFile(EntryFile)

	if err != nil {
		panic(err)
	}

	completeList := string(file)
	fmt.Printf("Complete File: %s", completeList)

	listArray := strings.Split(completeList, "\n")
	listArray = listArray[:len(listArray)-1]
	fmt.Printf("List Array: %s\n", listArray)
	for _, value := range listArray {
		fmt.Printf("value: %s", value)
		splitList := strings.Split(value, "   ")
		fmt.Printf(" l1: %s, l2: %s\n", splitList[0], splitList[1])
		l1, err := strconv.Atoi(splitList[0])
		if err != nil {
			panic(err)
		}
		heap.Push(ListOne, l1)
		l2, err := strconv.Atoi(splitList[1])
		if err != nil {
			panic(err)
		}
		heap.Push(ListTwo, l2)
	}

	//fmt.Printf("List one length: %d", ListOne.Len())
	for ListOne.Len() > 0 {
		l1 := heap.Pop(ListOne)
		l2 := heap.Pop(ListTwo)
		dist := math.Abs(float64(l1.(int) - l2.(int)))
		fmt.Printf("L1: %f, L2: %f, Dist: %f\n", l1, l2, dist)
		TotalDist += dist
	}

	fmt.Printf("Total Dist: %f\n", TotalDist)

}
