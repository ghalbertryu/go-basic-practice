package utils

import "fmt"

func PrintDividingLine(divName string) {
	fmt.Printf("\n=========== %s ===========\n", divName)
}

func PrintSlice(s any) {
	switch s.(type) {
	case []int:
		t := s.([]int)
		fmt.Printf("len=%d cap=%d %v\n", len(t), cap(t), s)
	case []uint8:
		t := s.([]uint8)
		fmt.Printf("len=%d cap=%d %v\n", len(t), cap(t), s)
	case [][]uint8:
		t := s.([][]uint8)
		fmt.Printf("len=%d cap=%d %v\n", len(t), cap(t), s)
	}
}
