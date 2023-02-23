package main

import (
	"fmt"
	"go-basic-practice/utils"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	mapCreate()
	mapMutate()
}

func mapCreate() {
	utils.PrintDividingLine("mapCreate")

	var m map[string]Vertex
	fmt.Printf("m=%v, is nil=%v\n", m, m == nil)

	m = make(map[string]Vertex)
	fmt.Printf("m=%v, is nil=%v\n", m, m == nil)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Printf("m=%v, is nil=%v\n", m, m == nil)
	fmt.Println(m["Bell Labs"])

	var m2 = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Printf("m2=%v, is nil=%v\n", m2, m2 == nil)
}

func mapMutate() {
	utils.PrintDividingLine("mapMutate")

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
