package main

import (
	"fmt"
	
)

// const(
// 	errorList = iota+ 6
// 	catSpecialist
// 	dogSpecialist
// 	snakeSpecialist
// )
// type Doctor struct{
// 	id int
// 	actorName string
// 	epissodes []string
// 	companions []string
// }
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

	// a := []int{}

	// fmt.Println(a)
	// a = append(a, 5)
	// fmt.Println(a)
	// a = append(a, 7,8,9,10)
	// fmt.Printf("%v\n", len(a))
	// fmt.Printf("%v\n", cap(a))
	// fmt.Printf("%v", a)

	//  var num float64 = 5
	// var result  = math.Sqrt(num)

	// fmt.Printf("%T", result)

	// const h bool = false

	// fmt.Println(h)


	// var matrix [][]int = [][]int{[]int{1,2,3},[]int{3,2,1},[]int{2,1,3}  }
	// fmt.Println(matrix[0])

	// a := []int{1, 2,3,4,5}
	// b := append(a[1:], 6)

	// fmt.Println(b)
	// statePopulation := make(map[string]int)
	// statePopulation = map[string]int{
	// 	"egypt" : 15000, 
	// 	"KSA" : 20000,
	// 	"UAE" : 17000,
	// } 
	// statePopulation["cairo"] = 10000
	// delete(statePopulation, "cairo")
	// fmt.Println(statePopulation)

	aDoctor := Doctor{
		 id : 2,
		 actorName:  "abdallah elassal",
		 companions: []string{
			"alaa ahmed",
			"ahmed",
		},
		epissodes: []string{
			"islam",
		},
	}
	fmt.Println(aDoctor)


}
