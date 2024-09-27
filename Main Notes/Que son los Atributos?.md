Tags: [Bases de Datos](../Indexes/Bases%20de%20Datos.md) Date: 2024-08-20 13:05
Status: #complete

# Que son los atributos?

Los _Atributos_ describen las caracteristicas que tendran nuestras _Entidades_.
Como recordaremos, las entidades no es mas que una representacion de un objeto o
concepto en la vida real, asi que, las entidades tendran que tener
obligatoriamente atributos que los describan.

Por ejemplo, En la entidad _Casa_ tendra atributos como _Direccion, Antiguedad,
Color, Puertas, Ventanas, Tamaño_, estos son los atributos que describen a la
entidad.

## Clasificacion de Atributos

Los atributos se clasifican en varios tipos:

- **Atributos Simples**: Son los datos o valores unicos de nuestra entidad que
  no puden divisibles en multiples datos. Por ejemplo, _Nombre y Edad_ son
  atributos simples en una Entidad _Persona_.
- **Atributos Compuestos**: Son los atributos que estan compuestos por multiples
  valores. Por ejemplo, el atributo _Direccion_ en una Entidad _Casa_ puede
  tener multiples valores como _Calle, Ciudad y Codigo Postal_
- **Atributos Clave**: Es el atributo que es _Clave Primaria_ en nuestra Entidad
  y que identifica de manera unica a la Entidad y la diferencia de otras. Por
  ejemplo, _Matricula_ en una Entidad _Estudiante_.
- **Atributos Multivalor**: Son los atributos que tienen mas de un valor en
  nuestra entidad. Por ejemplo, atributos como _Puertas, Ventanas_ en una
  Entidad _Casa_ son multivalor porque hay mas de una ventana y una puerta en
  cada Casa.
- **Atributos Derivados**: Son los atributos que nacen o derivan de otro
  atributo. Por ejemplo, _Antiguedad_ se deriva del atributo _Fecha de
  contruccion_ en una Entidad _Casa_.
- **Atributos Nulos**: Son los atributos que puden tener valores nulos a falta
  de informacion dispoinble en una Entidad.
- **Atributos Claves Extranjera**: Referencian a claves primarias de otras
  entidades y representan las relaciones entre entidades
