Tags: [Internet](../Indexes/Internet.md) Date: 2024-08-20 13:01 Status:
#complete

# Que es un nombre de dominio?

Un nombre de dominio es el encargado de identificar las webs o recursos en
internet por un nombre simple, por ejemplo "google.com".

Como ya se habiamos visto antes, el internet funciona por medio de direcciones
IPs, las cuales cada cliente y servidor tienen, no hay que confundirnos, un
nombre de dominio no es un IP, solo es el encargado de poder simplificar las IPs
a un texto facil de recordar.

Imaginemos por un momento que nuestro nombre es una IP, en este caso la IP seria
tu nombre completo con apellidos, supongamos "Jhon Doe Wish". La IP es el nombre
completo pero sabemos que nadie te llama por ese nombre, asi que las personas lo
abrevian a simplemente "Jhonny" asi que en este caso, "Jhonny" es tu nombre de
dominio.

No hay que confundirse. Las IPs no es una cadena de caracteres, sino una cadena
de numeros entre el 0 y el 255(por ejemplo `173.194.121.32`), es por eso que el
nombre de dominio simplifica la identificacion de las direcciones IP.

## DNS

Hablemos acerca del _DNS_, gracias a el existen los nombres de dominio. El
Sistema de nombres de dominio(DNS) es el encargado de traducir un nombre de
dominio a una direccion IP, cada vez que nosotros queremos acceder a Google,
escribimos en el navegador(cliente) "google.com" para acceder, el DNS traduce
"google.com" a una direccion IP(por ejemplo `190.124.12.41`).

Ahora que el el DNS a traducido nuestro nombre de dominio a una IP, nuestro
dispositivos pueden localizar el sitio web o recurso.

## TLDs

Alguna vez te has preguntado, por que los nombres de domino tinenen un ".com",
".net" o ".org" al final?. Estos son los TLD(_Dominio de nivel superior_).

Cada nombre de dominio se conforma por 3 partes(TLD, 2LD y 3LD) separados del
".", y se leen de derecha a izquierda. Cuando queremos acceder a paginas como
"google.com" , ".com" seria nuestro TLD y "google" el 2LD.

```
google.com
2DL   | TDL

// pero para el nombre de domino de Google UK
google.co.uk
 3LD |2LD|TLD

*En este caso 2LD indica el tipo de organizacion que registro el nombre(.co en el Reino Unido es para sitios registrados por empresas)*
```
