Tags: [Internet](../Indexes/Internet.md) Date: 2024-08-20 12:43 Status:
#complete

# Como funciona el internet?

Como ya sabemos, el internet es una red de dispositivos conectados entre si que
se comunican a traves de protocolos estandarizados. Estos protocolos definen
como se intercambia la informacion entre los dispositivos para un optimo
funcionamiento.

Cuando envias datos por internet, estos se dividen en paquetes que son
interceptados por una red global de routers que conforman el nucleo de internet.
El router es el encargado de examinar los paquetes recibidos y lo envia a un
siguiente router hasta llegar a su destino.

**Que es un protocolo?**: Un protocolo es un conjunto de reglas y normas que
definen como se intercambia informacion entre dispositivos y sistemas.

## Protocolos IP y TCP

Internet utiliza una gran variedad de protocolos estandarizados para asegurar el
manejo de los paquetes. El Protocolo de Informacion **(IP)** es el responsable
de enrutar los paquetes a su destino. El Protocolo de Controlo de Transmisiones
**(TCP)** asegura la correcta transmision de los paquetes de forma fiable y en
orden.

Estos son los protocolos basicos, pero hay una amplia variedad de protocolos,
algunos son:

- Sistema de Nombres de Dominio **(DNS)**
- Protocolo de Transferencia de Hipertexto **(HTTP)**
- Protocolo de Seguridad de Capas de Sockets Seguro/Transporte **(SSL/TLS)**

## Conceptos básicos

- **Paquete**: Una minima unidad de datos que se transmite a través de internet.
- **Router**: Un dispositivo que dirige paquetes de datos entre diferentes
  redes.
- **Direccion IP**: Un identificador único asignado a cada dispositivo en una
  red, utilizado para la ruta de los datos al destino correcto.
- **Nombre de dominio**: Un nombre legible por humanos que se utiliza para
  identificar un sitio web, como google.com.
- **DNS**: El sistema _Domain Name_ es responsable de traducir nombres de
  dominio en direcciones IP
- **HTTP**: El protocolo de transferencia de hipertexto se utiliza para
  transferir datos entre un cliente (como un navegador web) y un servidor (como
  un sitio web)
- **HTTPS**: Una version encriptada de HTTP que se utiliza para proporcionar
  comunicacion segura entre un cliente y un servidor
- **SSL/TLS**: Los protocolos _Secure Sockets Layer_ y _Transport Layer Security
  se utilizan para proporcionar una comunicacion segura en internet_

## Direcciones IP y Nombres de Dominio

Una direccion ip es un identificador unico asignado a cada dispositivo en una
red. Las direcciones ip de utilizan para enviar datos al destino correcto.

Las direcciones ip estan conformadas en una serie de cuatro numeros separados
por periodos, como "192.168.1.1"

**Entonces, que son los nombres de dominio?** Los nombres de dominio son nombres
legibles que sirven para identificar sitios web y otros recursos de internet.
Normalmente se componen de dos partes, por ejemplo: "google.com".

Pero el internet no entiende direcciones como "google.com" o "facebook.com", el
dominio solo sirve para eso, identificar por un nombre a un recurso en la web.

**Y entonces, si la web no entiende nombres, como podemos acceder a las webs
como google, amazon, etc?** Para eso se utiliza el sistema _Domain Name_ , como
ya vimos, el DNS es el encargado de traducir nombres como "google.com" a
direcciones ip para que asi podamos conectarnos a los sitios web o recursos.

DNS es una parte crucial en la infrestructura de internet. Al introducir un
nombre de dominio en un navegador web, el cliente envia una consulta DNS a un
servidor DNS, que devuelve la direccion IP correspondiente. Entonce el cliente
utiliza es direccion IP para conectarse al sitio web o recurso.

# References

[How does the internet work?](https://cs.fyi/guide/how-does-internet-work)
[The internet Explained](https://www.vox.com/2014/6/16/18076282/the-internet)
[How Does the Internet Work?](http://web.stanford.edu/class/msande91si/www-spr04/readings/week1/InternetWhitepaper.htm)
