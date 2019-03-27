package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubblingSort(sort []int) {
	for i := 0; i < len(sort); i++ {
		for j := i; j < len(sort); j++ {
			if sort[i] < sort[j] {
				tem := sort[i]
				sort[i] = sort[j]
				sort[j] = tem
			}
		}
	}
}

func selectSort(sort []int) {
	for i := 0; i < len(sort); i++ {
		min := sort[i]
		for j := i + 1; j < len(sort); j++ {
			if sort[j] > min {
				t := sort[j]
				sort[j] = min
				min = t
			}
		}
		sort[i] = min
	}
}

func insertSort(sort []int) {
	for i := 0; i < len(sort)-1; i++ {
		index := 0
		value := sort[i+1]
		for j := i + 1; j > 0; j-- {
			sort[j] = sort[j-1]
			if sort[j-1] > value {
				index = j
				break
			}
		}
		if index > 0 && sort[index-1] < value {
			sort[index-1] = value
		} else {
			sort[index] = value
		}
	}
}

func halfSort(sort []int, begin int, end int) {

	if begin == end {
		return
	}
	key := sort[begin]
	i := begin
	j := end
	for {
		if i == j {
			break
		}
		for ; j > i; j-- {
			if sort[j] > key {
				sort[i] = sort[j]
				sort[j] = key
				break
			}
		}
		if i == j {
			break
		}
		for i++; i < j; i++ {
			if sort[i] < key {
				sort[j] = sort[i]
				sort[i] = key
				break
			}
		}
	}
	if begin < j {
		halfSort(sort, begin, j)
	}
	if j < end {
		halfSort(sort, j+1, end)
	}
}

func main() {
	size := 100000
	var sort = make([]int, size)
	for i := 0; i < size; i++ {
		sort[i] = rand.Intn(99999)
	}
	fmt.Println("begin....")
	begin := time.Now().Unix()
	//halfSort(sort, 0, len(sort)-1) //99496
	bubblingSort(sort) //99496
	fmt.Printf("time:%d   ,%d \n", time.Now().Unix()-begin, sort[500])

}

func doEx(m func([]int), sort []int, name string) {
	fmt.Println("begin....")
	t := make([]int, len(sort))
	copy(t, sort)
	begin := time.Now().Unix()
	m(t)
	fmt.Printf("time:%d,%v,%d \n", time.Now().Unix()-begin, name, t[500])
}

func test() {
	size := 100000
	var sort = make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		sort[i] = rand.Intn(99999)
	}
	go doEx(bubblingSort, sort, "bubblingSort")
	go doEx(selectSort, sort, "selectSort")
	go doEx(insertSort, sort, "insertSort")
	time.Sleep(time.Hour)

}
