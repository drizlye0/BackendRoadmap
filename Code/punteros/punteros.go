package main

import "fmt"

func main() {
	k :=  100
	p := &k
	*p = 20

	fmt.Println("Con * accedo al valor que contiene la direccion en memoria de lo que apunto")
	fmt.Println(*p)

	fmt.Println("Con & accedo ala direccion en memoria de mi puntero, no la que tiene la variable a la que apunto")
	fmt.Println("puntero: ", &p)
	fmt.Println("variable a la que apunto", &k)

	fmt.Println("Si quiero la direccion a la que apunto con el puntero, simplemente uso p")
	fmt.Println(p)
}
