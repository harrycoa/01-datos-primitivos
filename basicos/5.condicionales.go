package main

import "fmt"

/*
== igual
!= diferente
< menor
> mayor
<= menor  o igual
>= mayor o igual
&& and
|| or
*/


func main()  {

	x := 10
	y := 12
	if (x >= y) {
		fmt.Printf("%d es mayor que %d \n",x,y)
	} else if x < y{
		fmt.Printf("%d es mayor que %d \n",y,x)
	} else {
		fmt.Println("son iguales ")
	}
}
