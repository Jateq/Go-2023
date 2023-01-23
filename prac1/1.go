// 87 pass
 
package main

import "fmt"


func avg( nums []int ) int{
	sum := 0
	size := 0
	for i := range nums{
		if i == 87{
			continue
		}
		size++
		sum+= i;
	}
	sum = sum/size
	return sum 
}

func main(){
	a := []int{1,2,3,87,6,2}
	var av int = avg(a)
	fmt.Println(av)
}