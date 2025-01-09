package main

import "fmt"

// const(
// 	errorList = iota+ 6
// 	catSpecialist 
// 	dogSpecialist
// 	snakeSpecialist
// )
 
func main(){
	//i := false
	// i  := 1==2 
	// fmt.Printf("%v , %T\n" ,i , i)

	// s := 10
	// fmt.Println(s)
	// a := 8
	// fmt.Println( a << 3 )
	// fmt.Println( a >> 3 )
	// var a int = 8
	// var b int8 = 10
	// fmt.Println(a + int(b))

	// var orderList int = dogSpecialist
	//fmt.Printf("%v",  catSpecialist)

	a := []int{}

	fmt.Println(a)
	a = append(a, 5)
	fmt.Println(a)
	a = append(a, 7,8,9,10)
	fmt.Printf("%v\n", len(a))
	fmt.Printf("%v\n", cap(a))
	fmt.Printf("%v", a)


}
