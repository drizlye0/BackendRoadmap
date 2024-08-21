Tags: [Bases de Datos](Bases%20de%20Datos.md)
Date: 2024-08-20 13:05
Status: #complete 

# Que son los atributos?

Los *Atributos* describen las caracteristicas que tendran nuestras *Entidades*. Como recordaremos, las entidades no es mas que una representacion de un objeto o concepto en la vida real, asi que, las entidades tendran que tener obligatoriamente atributos que los describan. 

Por ejemplo, En la entidad *Casa* tendra atributos como *Direccion, Antiguedad, Color, Puertas, Ventanas, Tamaño*, estos son los atributos que describen a la entidad.

## Clasificacion de Atributos

Los atributos se clasifican en varios tipos:
- __Atributos Simples__: Son los datos o valores unicos de nuestra entidad que no puden divisibles en multiples datos. Por ejemplo, *Nombre y Edad* son atributos simples en una Entidad *Persona*.
- __Atributos Compuestos__: Son los atributos que estan compuestos por multiples valores. Por ejemplo, el atributo *Direccion* en una Entidad *Casa* puede tener multiples valores como *Calle, Ciudad y Codigo Postal* 
- __Atributos Clave__: Es el atributo que es *Clave Primaria* en nuestra Entidad y que identifica de manera unica a la Entidad y la diferencia de otras. Por ejemplo, *Matricula* en una Entidad *Estudiante*.
- __Atributos Multivalor__: Son los atributos que tienen mas de un valor en nuestra entidad. Por ejemplo, atributos como *Puertas, Ventanas* en una Entidad *Casa* son multivalor porque hay mas de una ventana y una puerta en cada Casa.
- __Atributos Derivados__: Son los atributos que nacen o derivan de otro atributo. Por ejemplo, *Antiguedad* se deriva del atributo *Fecha de contruccion* en una Entidad *Casa*.
- __Atributos Nulos__: Son los atributos que puden tener valores nulos a falta de informacion dispoinble en una Entidad.
- __Atributos Claves Extranjera__: Referencian a claves primarias de otras entidades y representan las relaciones entre entidades