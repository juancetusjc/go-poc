package main
import (
	//"net/http"
	"fmt"
)


func main(){
	fmt.Println("inicio")	
	var result int
	result= operacionesMidd(sumar)(2,3)
	fmt.Println(result)
	result= operacionesMidd(restar)(3,2)
	fmt.Println(result)
	result= operacionesMidd(multiplicar)(2,3)
	fmt.Println(result)
}

func operacionesMidd( f func(int,int) int) func(int, int) int{
 return func(x,y int) int{
	 fmt.Println("Inicio de Operacion")
	 return f(x,y)
 }
}


func sumar(a,b int) int{
	return a+b	
}

func restar(a,b int) int{
	return a-b	
}

func multiplicar(a,b int) int{
	return a*b	
}