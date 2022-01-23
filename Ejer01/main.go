package main 

import (
	"fmt"
	"strings"
	"time"
)


var numero int


func main(){
	/*go miNombreLento("juan")
	fmt.Println("Etoy aqui")
	var x string
	fmt.Scanln(&x)*/

	canal1:= make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegue hasta aqui")
	
	msg :=<- canal1
	fmt.Println(msg)

}


func bucle(canal chan time.Duration){
	inicio:= time.Now()
	sum  := 1
	for i :=0; i<100000000000; i++{
		sum = sum+i
	}
	fmt.Println("suma:")
	fmt.Println(sum)
	final := time.Now()
	canal<- final.Sub(inicio)	
}

func miNombreLento(nombre string){
	letras:=strings.Split(nombre,"")
	for _,letra:= range letras{
		time.Sleep(1000*time.Millisecond)
		fmt.Println(letra)
	}
} 

