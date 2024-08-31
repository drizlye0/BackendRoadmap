Tags: [Internet](../Indexes/Internet.md)
Date: 2024-08-20 12:59
Status: #complete 

# Como funciona un Navegador?

Un navegadro web tiene el principal objetivo de poder entregarte contenido del internet en en una ventana. Esta contenido pude ser HTML,PDF o incluso videos e imagenes.

En la actualidad, existen un amplio numero de navegadores web, y estos ser rigen de un estandar que se ha establecido por lo largo del tiempo.

Un navegador web consta de multiples partes:
- __Interfaz de usuario__: Es todo lo que el usuario ve y usa, es la cara inicial del navegador y muy importante para el usuario.
- __Motor del Navegador__: Ordena las acciones entre la UI y el motor de renderizacion 
- __Sistema de Renderizado__: Es el encargado de poder renderizar los archivos HTML y CSS.
- __Herramienta de Redes__:  Se refiere a los auxiliares para el correcto funcionamiento del recibimiento y envio de datos.
- __Javascript Interpreter__: El interpretador para poder ejecutar codigo Javascript
- __Persistencia de Datos__: Los navegadores cuentan con una capa de persistencia para poder almacenar datos de manera local(por ejemplo las cookies) 
- __UI Backend__: Es el funcionamiento que hay por detras para que UI del navegador pueda dibujar widgets, ventanas, etc.

## Funcionamiento

El funcionamiento de un navegador inicia desde que un usuario busca un pagina web, para poder explicar todo esto usare como usuario a Bobby. 

Bobby quiere entrar a facebook para ver lo que han compartido sus amigos y contestar mensajes, asi que, Bobby entra a su navegador y escribre en la barra de busqueda "facebook.com". Lo que acaba de escribir Bobby(sin darse cuenta) es una URL, lo que significa que dentro de esta URL viene un nombre de dominio, una peticion HTTP con sus respectivos headers.
Lo primero que hace nuestro navegador es pedirle a un servidor DNS, que traduzca el nombre de dominio "facebook.com" a una IP que la laptop de Bobby pueda buscar. Una vez que el servidor DNS pueda encontrar la IP que le corresponse a el nombre de dominio, el servidor responde y entrega la IP a la laptop de Bobby.

Entonces una vez que la laptop obtiene la IP, pasaran dos cosas:
1. La laptop de Bobby guardara la IP del nombre del dominio en cache para que la proxima vez que Bobby quiera entrar a facebook, la laptop no tenga que volver a pedir la IP a un servidor DNS. A esto se le conoce como DNS Cache.
2. La lapotop de Bobby empieza a buscar el servidor con la direccion IP que recibio del servidor DNS para poder hacerle la peticion HTTP. 

Una vez que sea encontrado el servidor y hecha la peticion, el servidor respondera con archivos HTML,CSS,JS que el navegador tendra que renderizar para que el Bobby pueda verlos.