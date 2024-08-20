Tags: [[Internet]]
Date: 2024-08-20 13:01
Status: #complete 

# Servidor Web

Un servidor web no es mas que una computadora(hardware) que almacena archivos de paginas web(archivos html, css, js, imagenes, etc)  y que tiene abierto un software de servidor web.

El software de servidor web se el encargado de servir estos archivos a los usuarios que accedan a nuestra sitio web. Este servidor es HTTP, es decir, puede comprender URLS para obtener paginas web gracias a HTTP(el protocolo que usan los navegadores para obtener paginas web). Es muy importante comprender como funciona las URLS y los DNS para entender los servidores HTTP ya que todo esta relacionado. 

Los servidores web estan divididos en dos grupos:
- __Servidores estaticos__: Los servidores estaticos son solo capazes de servir contenido estatico y no modificado por el servidor hacia el navegador(por ejemplo, un articulo cientifico o un blog personal).
- __Servidores dinamicos__: La principal diferencia es que estos servidores actualizan su contenido usando una base de datos antes de servirlos al navegador(por ejemplo, sitios como amazon o aliexpress se actualizan cada dia).

De una manera mas simple, cuando buscas algo por internet, ya sea una pagina, sitio, imagen, pdf, etc. El navegador mandara una o varias peticiones HTTP al servidor para obtener lo que deseas, una vez que el servidor(hardware) haya recibido la peticion, el software del servidor podra mandar una respuesta al navegador. 

# Hosting y VPS

A simples rasgos. Un hosting no es mas que un servicio por el cual pagas para que alojen tu pagina o sitio web en un servidor remoto compartido por otros usuarios.

Definamos un poco mas. El hosting nos ayuda a que nuestro sitio web o pagina este en disponible para los usuarios todo el tiempo que sea posible, en este caso cuando hablamos de hosting, nos referimos a un hosting compartido por otros usuarios que tambien quieran alojar sus sitios webs. Esto es bueno si sitio o pagina web es muy poca visitada o es un proyecto no muy grande. 

Las deventajas de un hosting compartido, es que al ser compartido por otros usuarios, las limitaciones estan muy presentes, ya que no puedes configurar libremente el servidor, lo que hace que en proyectos grandes, esten destinados a tener caidas o problemas de rendimiento.

## VPS

Los VPS *(Virtual Partition Server)* es una forma de alojar nuestro sitio web, sigue siendo un hosting compartido con demas usuarios, pero cada usuario tiene un parte virtual del servidor que puede configurar a su gusto.

A diferencia de un hosting compartido, un VPS es una forma de rentar una parte de un servidor, en concreto, rentamos un hardware que nosotros podemos pagar por los recursos que vayamos a usar, por ejemplo, si nuestro proyecto es un ecommerce que recibe le promedio de solicitudes HTTP, podriamos tener un VPS que tenga unos recursos no tan altos(4GB RAM, 150GB de almacenamiento, etc). 

La principal ventaja de los VPS es que son muy faciles de administrar, si un dia demandante de compras(como un black friday, cybermonday) nuestro ecommerce empieza a tener problemas de rendimiento, podemos ampliar los recursos de nuestro VPS, ya que la mayoria cuenta con un panel de control para sus usuarios, y podemos solamente ampliar nuestros recursos para esa ocasion y despues bajar los recursos si ya no se necesitan.

## Servidor Dedicado

Como su nombre lo dice, un servidor dedicado para poder hacerle lo que nosotros deseemos. Esta opcion es mas acertada para empresas que requiran una infrestructura muy grande para proyectos.

Un servidor dedicado presenta ciertas desventajas:
- __Mantenibilidad__: Al ser un servidor dedicado, nosotros nos hacemos responsables de mantenerlo en funcionamiento en lugar de una empresa dedicada a hostings.
- __Escabilidad__: Un servidor dedicado es escalable pero los precios para escalarlo ya no son iguales a los de un VPS, y antes de escalarlo se debe de estar seguro si es necesario o habremos hecho una inversion injustificada.