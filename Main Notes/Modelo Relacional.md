Tags: [[Bases de Datos]]
Date: 2024-08-20 13:03
Status: #complete 

# DB Relacionales: Modelo Relacional

El modelo relacional representa como se guarda la informacion en una base de datos relacional. En una base de datos relacional la informacion se guarda en columnas y filas, que crean relaciones y esas relaciones se guardan en tablas.

En el modelo relacional, cada tabla se le asigna un nombre unico. Ejemplo:

__Tabla STUDENT__:  

| ID  | Name   | Address | Phone  | Age |
| --- | ------ | ------- | ------ | --- |
| 1   | RAM    | DELHI   | 890192 | 18  |
| 2   | RAMESH | GURGAON | 891029 | 18  |
| 3   | SUJIT  | ROHTAK  | 892012 | 20  |
| 4   | SURESH | DELHI   | 891204 | 18  |

## Integridad clave

Cada relacion en las bases de datos relacionales debe tener un atributo el cual sea unico para cada tupla, a esto se le conoce como un clave o *key*. Por ejemplo; *ID* en la relacion *STUDENT*.

Asi que una clave o *key* tiene dos propiedades:
- Debe ser unico para todas las tuplas
- NO puede tener valores NULL

## Caracteristicas del modelo relacional

- Los datos estan representados en filas y columnas llamadas relaciones
- Los datos se almacenan en tablas que tienen relaciones entre ellas llamadas modelo de relacion
- Cada fila representa una sola entidad
- Cada columna tiene un nombre distinto y representan atributos