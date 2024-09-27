Tags: [Bases de Datos](../Indexes/Bases%20de%20Datos.md) Date: 2024-08-20 13:04
Status: #complete

# Que es una Entidad?

Una entidad es una forma de representar un objeto, evento, persona, concepto de
la vida real. La entidad es clave para el modelo _Entidad-Relacion_. La entidad
tiene un conjunto de _Atributos_ que describen las caracteristicas de la
entidad.

Por ejemplo, una instancia de la entidad _Persona_ **Tabla Persona**:

| Nombre | Apellidos |
| ------ | --------- |
| Jhon   | Doe       |

_Jhon Doe_ es una entidad de _Persona_ que a su vez tiene atributos _Nombre_ y
_Apellidos_

## Clasificacion de Entidades

Existen _Entidades Fuertes_ y _Entidades Debiles_, se pueden diferenciar de una
manera sencilla, ya que las entidades debiles carecen de una _Clave Primaria_ y
su \_\_existencia depende de una Entidad Propietaria.

Una entidad debil tambien tiene la caracteristica de que siempre tendra una
_Relacion de 1 a Muchos_
