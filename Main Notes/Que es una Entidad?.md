Tags: [Bases de Datos](Bases%20de%20Datos.md)
Date: 2024-08-20 13:04
Status: #complete 

# Que es una Entidad?

Una entidad es una forma de representar un objeto, evento, persona, concepto de la vida real. La entidad es clave para el modelo *Entidad-Relacion*. La entidad tiene un conjunto de *Atributos* que describen las caracteristicas de la entidad.

Por ejemplo, una instancia de la entidad *Persona*
__Tabla Persona__:

| Nombre | Apellidos |
| ------ | --------- |
| Jhon   | Doe       |
*Jhon Doe* es una entidad de *Persona* que a su vez tiene atributos *Nombre* y *Apellidos*

## Clasificacion de Entidades
Existen *Entidades Fuertes* y *Entidades Debiles*, se pueden diferenciar de una manera sencilla, ya que las entidades debiles carecen de una *Clave Primaria* y su __existencia depende de una Entidad Propietaria.

Una entidad debil tambien tiene la caracteristica de que siempre tendra una *Relacion de 1 a Muchos*