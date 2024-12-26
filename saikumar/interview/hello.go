package main

import "fmt"

func main() {

	// for i := 5; i >=1; i-- {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print(j," ")
	// 	}
	// 	fmt.Println("")
	// }

	// 1
	// 5 6
	// 9 7 4
	// 8 0 3 2
	fmt.Println("--------------------------------")
	s := []int{1, 5, 6, 9, 7, 4, 8, 0, 3, 2}
	// i := 0
	// j := i
	for i :=0; i < 10; i++ {
		
		if s[i]==1{
			fmt.Println()
		}
		if s[i]==5{
			fmt.Println()
		}
		if s[i]==9{
			fmt.Println()
		}
		if s[i]==8{
			fmt.Println()
		}
		fmt.Print(s[i]," ")
		
		// for j = i; j <= i; j += 1 {
		// 	fmt.Println(i, j)
		// 	// fmt.Print(s[i])
		// }

	}

}
