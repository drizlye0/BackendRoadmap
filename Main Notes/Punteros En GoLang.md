Tags: [GoLang](../Indexes/GoLang.md) 
Date: 2024-08-21 19:02
Status: #progress 
# Punteros En GoLang
Los punteros en Go son fundamentales para poder trabajar con el lenguaje, pero que es un puntero?

## Definicion
Cuando se habla de punteros nos referimos a un tipo especial de variable que en lugar de guardar un valor como un entero, string o float, guarda una *direccion en memoria*.

Cuando nosotros creamos una variable, __estamos creando un espacio en la memoria de nuestro programa__, en este espacio de memoria guardamos el valor que le asignamos a nuestra variable, __pero tambien guardamos una direccion__.

Para entenderlo de mejor forma, hagamos una representacion grafica de nuestra memoria:

__Memoria__

| Nombre     | Direccion    | Valor        | Tipo       |
| ---------- | ------------ | ------------ | ---------- |
| myVar1<br> | 0xc000056040 | 100          | int        |
| myVar2     | 0xc000012130 | "Hola mundo" | string<br> |
| myPointer  | 0xc000034072 | 0xc000012130 | *string    |

En esta representacion grafica hay tres espacios de memoria, cada uno de ellos tiene *nombre, direccion, valor y tipo*. En la variable *myVar2* se guarda un *"Hola mundo"* que es de tipo string y tiene una direccion que es *0xc000012130*, y tambien se encuentra un puntero *myPointer* que guarda la direccion de *myVar* como valor y que apunta a un tipo int.

### Simbolos * &
Con estos simbolos podemos hacer dos cosas con los punteros:
- Con el * podemos acceder al valor de la variable a la que apuntamos.
- Con & podemos tener la direccion en memoria de nuestro puntero, que es diferente a la variable a la que apuntamos.


## Declaracion y uso de los punteros en GoLang
Archivo con ejemplos y usos de los punteros:
- [punteros](../Code/punteros/punteros.go)

# Referencias
[Que es un puntero?](https://lenguajedeprogramacion.com/programacion-c/que-es-un-puntero-usos/)
[Entendiendo punteros en programacion con golang](https://www.youtube.com/watch?v=gvjON1S0drk)
