package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	intTest()
	stringTest()
	pringFTest()
}

func intTest() {
	fmt.Println("intTest\n=======")

	var a int8 = 3
	var b int32 = 3
	var c int64 = 3
	var d int = 3
	fmt.Println(a, b, c)
	fmt.Println(unsafe.Sizeof(a),
		unsafe.Sizeof(b),
		unsafe.Sizeof(c),
		unsafe.Sizeof(d))

	fmt.Println("=======\n")
}

func stringTest() {
	fmt.Println("stringTest\n=======")

	str := `
		package main
		
		import (
			"fmt"
			"strconv"
			"unsafe"
		)
	`
	fmt.Println(str)

	fmt.Println("=======\n")
}

func pringFTest() {
	fmt.Println("pringFTest\n=======")

	numStr := "12345"
	fmt.Printf("digit:%d, char:%c, string:%s, type:%T, value:%v \n",
		numStr, numStr, numStr, numStr, numStr)
	num, _ := strconv.ParseInt(numStr, 10, 64)
	fmt.Printf("digit:%d, char:%c, string:%s, type:%T, value:%v \n",
		num, num, num, num, num)

	fmt.Println("=======\n")
}
