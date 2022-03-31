package main

import "fmt"

var total float64
var count int

func calTotal(value []float64) (float64, int) {
	var totalVal float64
	var inCount int
	for i := 0; i < len(value); i++ {

		totalVal += value[i]
		inCount = len(value)
	}
	return totalVal, inCount
}

func main() {
	//slices template
	room1 := []float64{20, 21, 23, 25, 22}
	room2 := []float64{27, 23, 25, 20, 30, 24}
	room3 := []float64{22, 23, 24, 22}

	//Insert code here

	// fmt.Printf("%v, %T", room1[1], room1[1])

	sliceList := [][]float64{room1, room2, room3}

	// fmt.Println(len(sliceList))

	for i := 0; i < len(sliceList); i++ {
		sVal, sCount := calTotal(sliceList[i])
		total += sVal
		count += sCount

	}
	average := total / float64(count)

	fmt.Printf("%.02f", average)

}
