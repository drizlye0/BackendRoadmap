Tags: [Internet](../Indexes/Internet.md) Date: 2024-08-20 13:01 Status:
#complete

# Que es HTTP?

HTTP es el protocolo para transferir datos entre un cliente(como un navegador
web) y un servidor(como un sitio web).

Es sencillo, por ejemplo, usando el cliente quiere acceder a una pagina en
internet, el cliente envia una solicitud HTTP al servidor donde se encuentra la
pagina a la que quiere acceder, si la solicitud es correcta, el servidor envia
una respuesta HTTP al cliente, conteniendo la pagina que solicitaba el cliente.

HTPP en si depende de TCP/IP para obtener solicitudes y respuestas entre el
cliente y el servidor. A diferencia de TCP que usa el puerto 80 por defecto,
HTTPS utiliza el puerto 443

La primera version de HTTP fue **HTTP/0.9** que se presento en 1991, esta
primera version solo contaba con un solo **metodo** llamado **GET**. El metodo
GET es el encarga de servir contenido de un servidor a un cliente.

Aqui un ejemplo de como se comunican el cliente y el servidor mediante
solicitudes HTTP/0.9

```
// solicitamos al servidor index.html a traves del metodo GET
GET /index.html

// recibimos una respuesta por parte del servidor
(response body)
(connection closed)
```

Actualmente el protocolo HTTP a evolucionado a lo largo del tiempo y a
introducido cosas como metodos POST, PUT, DELETE, PATCH, HEAD, etc.

## HTTP/2 - 2015

Antes de hablar en HTTP/2 hay que conocer el transfondo de este. El protocolo
SPDY nace de la mano de google en 2009, que pretendia disminuir la latencia para
mejorar el rendimiento de la red, pero en 2015, Google no queria tener dos
estandares, asi que SPDY fue fusionado con HTTP y asi nace HTTP/2

Las principales diferencias entre HTTP y HTTP/2:

- Binario en lugar de texual
- Multiples solicitudes de HTTP asincronicas en una sola conexion
- Comprension de Heders mediante HPACK
- Server Push - Multiples respuestas para una sola solicitud
- Solicitud de priorizacion
- Seguridad
