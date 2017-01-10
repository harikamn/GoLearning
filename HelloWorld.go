package main
import "fmt"

func main(){
	fmt.Println("Hello World")
	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
	for i=1;i<=100;i++ {

		if i%3==0{
			fmt.Println(i)
		}

	}

	 a := [5]int{1,2,3,4,5}
	for i=0;i<len(a);i++{
		fmt.Println(a[i])
	}



}
